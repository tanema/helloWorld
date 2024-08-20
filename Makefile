.DEFAULT_GOAL := build
.PHONY: build

build:
	@GO111MODULE=on go build -o netlify/functions/hello .
