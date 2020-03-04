FROM golang:1.1.12

MAINTAINER An Do <andodevel@gmail.com>

ENV GOPATH /go
ENV GO111MODULE on

COPY . /go/src/github.com/andodevel/clock_server
WORKDIR /go/src/github.com/andodevel/clock_server
RUN make ci && make install

ENTRYPOINT ["/go/bin/clock_server"]
