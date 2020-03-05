# TODO: Correct this docker build
FROM golang:1.13-alpine
RUN apk --no-cache add bash ncurses make git gcc libtool musl-dev upx

LABEL maintainer="An Do <andodevel@gmail.com>"

ENV GOPATH /go
ENV GO111MODULE on
ENV PATH="${PATH}/bin:${PATH}"

COPY . /go/src/github.com/andodevel/clock_server
WORKDIR /go/src/github.com/andodevel/clock_server
RUN make install

ENTRYPOINT ["release"]
