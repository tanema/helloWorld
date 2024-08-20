.DEFAULT_GOAL := build
.PHONY: build

build:
	@go build -mod=mod -o netlify/functions/hello .
