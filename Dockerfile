FROM golang:1.1.12

MAINTAINER An Do <andodevel@gmail.com>

ENV GOPATH /go
ENV GO111MODULE on

COPY . /go/src/github.com/andodevel/go-echo-template
WORKDIR /go/src/github.com/andodevel/go-echo-template
RUN make ci && make install

ENTRYPOINT ["/go/bin/go-echo-template"]
