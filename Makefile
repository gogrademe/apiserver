NAME=GoGradeMeAPI
HARDWARE=$(shell uname -m)
VERSION=0.1.0

build:
	mkdir -p stage
	GOOS=linux go build -o stage/$(NAME)
	docker build -t gogrademe/api-server .

release:
	rm -rf release
	mkdir release
	GOOS=linux go build -o release/$(NAME)
	cd release && tar -zcf $(NAME)_$(VERSION)_linux_$(HARDWARE).tgz $(NAME)
	GOOS=darwin go build -o release/$(NAME)
	cd release && tar -zcf $(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz $(NAME)
	rm release/$(NAME)
	echo "$(VERSION)" > release/version
	echo "gogrademe/$(NAME)" > release/repo
	gh-release

.PHONY: release
