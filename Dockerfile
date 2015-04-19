FROM gliderlabs/alpine:3.1
EXPOSE 5005

COPY . /go/src/github.com/gogrademe/apiserver
RUN apk-install go git mercurial \
	&& cd /go/src/github.com/gogrademe/apiserver \
	&& export GOPATH=/go \
	&& export PATH=$PATH:/go/bin \
	&& go get github.com/tools/godep \
	&& godep restore \
	&& go build -o /bin/apiserver \
	&& rm -rf /go \
	&& apk del go git mercurial

CMD ["/bin/apiserver"]
