NAME=apiserver
VERSION=$(shell git describe --tags)

deps:
	go get -u github.com/tools/godep
	godep restore

build:
	go build -ldflags "-X main.version $(VERSION)" -o build/$(NAME)

.PHONY: build/$(NAME)
