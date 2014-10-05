FROM golang

ADD . /go/src/github.com/Lanciv/GoGradeAPI

WORKDIR /go/src/github.com/Lanciv/GoGradeAPI

RUN go env

RUN go get github.com/tools/godep

RUN godep restore ./...

RUN godep go build -o /usr/bin/GoGradeAPI main.go


EXPOSE 5005

ENTRYPOINT ["/usr/bin/GoGradeAPI"]
# CMD ["-staticDir=/opt/www"]
