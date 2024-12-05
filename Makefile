.PHONY: build deps

deps:
	npm i && \
	go mod tidy

build:
	go mod tidy && \
	templ generate && \
	tailwind -o public/static/main.css -i public/preset.css && \
	go build -o tmp/server cmd/server.go