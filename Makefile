run: locales build
	./gtd_bot

locales:
	cd i18n && ruby gen.rb | gofmt > i18n.go

build:
	go build -o gtd_bot ./main
