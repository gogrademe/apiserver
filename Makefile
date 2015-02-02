NAME=apiserver
VERSION=$(shell git describe --tags)

# HARDWARE=$(shell uname -m)
#
# build/container: stage/$(NAME) Dockerfile
# 	docker build -t $(NAME) .
# 	touch build/container
#
deps:
	go get github.com/tools/godep
	godep restore
	
build:
	go build -ldflags "-X main.version $(VERSION)" -o build/$(NAME)

#
# stage/$(NAME): build/$(NAME)
# 	mkdir -p stage
# 	cp build/$(NAME) stage/$(NAME)
#
# release: build/container
# 	rm -rf release
# 	mkdir release
# 	docker tag $(NAME) gogrademe/$(NAME)
# 	docker push gogrademe/$(NAME)
# 	# cd release && tar -zcf $(NAME)_$(VERSION)_linux_$(HARDWARE).tgz ../build/$(NAME)
# 	# echo "$(VERSION)" > release/version
# 	# echo "gogrademe/$(NAME)" > release/repo
# 	# gh-release
#
#
.PHONY: build/$(NAME)
# clean:
# 	rm -rf build
# 	rm -rf release
# 	rm -rf stage
