FROM lanciv/golang-base:1.3

ADD . /go/src/github.com/Lanciv/GoGradeAPI

WORKDIR /go/src/github.com/Lanciv/GoGradeAPI

RUN go get -u github.com/tools/godep

RUN go env

RUN godep restore ./...
RUN godep go clean ./...
RUN godep go build -o /usr/bin/GoGradeAPI main.go

# Clean all the unused packages

RUN apt-get autoremove -y
RUN apt-get clean all && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 5005

ENTRYPOINT ["/usr/bin/GoGradeAPI"]
# CMD ["-staticDir=/opt/www"]
