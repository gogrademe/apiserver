FROM alpine:3.1
ENTRYPOINT ["/bin/apiserver"]
EXPOSE 5005

COPY . /go/src/github.com/gogrademe/apiserver
RUN apk-install go git mercurial \
	&& cd /go/src/github.com/gogrademe/apiserver \
	&& export GOPATH=/go \
	&& export PATH=$PATH:/go/bin \
	&& go get github.com/tools/godep \
	&& godep restore \
	&& go build -ldflags "-X main.version $(git describe --tags)" -o /bin/apiserver \
	&& rm -rf /go \
	&& apk del go git mercurial
