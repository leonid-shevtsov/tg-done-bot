require 'json'

wunderlist = JSON.parse(File.read(ARGV[0]), symbolize_names: true)

inbox_id = wunderlist[:data][:lists].detect {|l| l[:title] == 'inbox'}[:id]
inbox_tasks = wunderlist[:data][:tasks].select{|t| t[:list_id] == inbox_id && t[:completed] == false}

wunderlist[:data][:notes].each do |note|
  task_id = note[:task_id]
  if task = inbox_tasks.detect{|t| t[:id] == task_id}
    task[:notes] ||= []
    task[:notes] << note[:content]
  end
end

def quote_string(s)
  s.gsub(/\\/, '\&\&').gsub(/'/, "''").gsub("\r", "").gsub("\n", "\\n")
end

inbox_tasks.each do |task|
  text_items = task[:notes] || []
  text_items.unshift(task[:title])
  text = text_items.join("\n")
  puts "INSERT INTO inbox_items (user_id, text, created_at) VALUES (#{ARGV[1]}, '#{quote_string(text)}', '#{task[:created_at]}');"
end
