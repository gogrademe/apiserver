FROM ubuntu:13.10

RUN apt-get update -q
RUN apt-get install -qy build-essential curl git mercurial bzr
RUN curl -s https://storage.googleapis.com/golang/go1.3beta2.src.tar.gz | tar -v -C /usr/local -xz
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1
ENV PATH /usr/local/go/bin:/opt/bin:$PATH
RUN mkdir /opt/bin
ENV GOPATH /opt

RUN go get github.com/tools/godep


ADD . $GOPATH/src/github.com/Lanciv/GoGradeAPI

RUN cd $GOPATH/src/github.com/Lanciv/GoGradeAPI && go get .
RUN cd $GOPATH/src/github.com/Lanciv/GoGradeAPI && go build .

# Clean all the unused packages
RUN apt-get autoremove -y
RUN apt-get clean all


CMD 'gogradeapi'
