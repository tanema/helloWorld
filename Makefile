.DEFAULT_GOAL := build
.PHONY: build

build:
	@go build -o netlify/functions/hello .
