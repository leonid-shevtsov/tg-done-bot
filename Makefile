run: locales
	go run main/main.go

locales:
	cd i18n && ruby gen.rb | gofmt > i18n.go
