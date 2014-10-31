NAME=apiserver
HARDWARE=$(shell uname -m)
VERSION=0.1.1

build/container: stage/$(NAME) Dockerfile
	docker build -t $(NAME) .
	touch build/container

build/$(NAME): *.go
	GOOS=linux GOARCH=amd64 go build -o build/$(NAME)

stage/$(NAME): build/$(NAME)
	mkdir -p stage
	cp build/$(NAME) stage/$(NAME)

release: build/container
	rm -rf release
	mkdir release
	docker tag $(NAME) gogrademe/$(NAME)
	docker push gogrademe/$(NAME)
	# cd release && tar -zcf $(NAME)_$(VERSION)_linux_$(HARDWARE).tgz ../build/$(NAME)
	# echo "$(VERSION)" > release/version
	# echo "gogrademe/$(NAME)" > release/repo
	# gh-release


.PHONY: clean release
clean:
	rm -rf build
	rm -rf release
	rm -rf stage
