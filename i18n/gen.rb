require 'yaml'
require 'active_support/core_ext/string'

data = YAML.load_file('en.yml')

puts <<EOT
package i18n

import "fmt"

func pluralize(count int, one string, other string) string {
  if (count % 10) == 1 {
    return fmt.Sprintf(one, count)
  } else {
    return fmt.Sprintf(other, count)
  }
}
EOT


def pluralized?(hash)
  hash.keys.sort == %w(one other)
end

def print_locale_type(hash, type_name)
  hash.each do |key, value|
    if value.is_a?(Hash) && !pluralized?(value)
      nested_type_name = "#{type_name}#{key.camelcase}"
      print_locale_type(value, nested_type_name)
    end
  end
  puts "type #{type_name} struct{"
  hash.each do |key, value|
    if value.is_a?(Hash)
      if pluralized?(value)
        puts "#{key.camelcase} func(count int) string"
      else
        puts "#{key.camelcase} #{type_name}#{key.camelcase}"
      end
    else
      puts "#{key.camelcase} string"
    end
  end
  puts "}"
  puts
end

def print_locale_value(hash, type_name, top_level = false)
  puts "#{type_name}{"
  hash.each do |key, value|
    if value.is_a? Hash
      if pluralized?(value)
        puts "func(count int) string { return pluralize(count, \"#{value["one"]}\", \"#{value["other"]}\")  },"
      else
        print_locale_value(value, "#{type_name}#{key.camelcase}")
      end
    else
      puts "\"#{value}\","
    end
  end
  puts "}#{top_level ? "" : ","}"
end

print_locale_type(data.values.first, "Locale")

data.each do |key, value|
  puts
  print "var #{key.camelcase} = "
  print_locale_value(value, "Locale", true)
end
