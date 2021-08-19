FROM golang:1.16.3-alpine3.13

WORKDIR /go/src/github.com/silvergama/unico
ADD . .

RUN apk add --no-cache pkgconfig gcc make musl-dev

# "make install" should be run before building image
RUN GOOS=linux GOARCH=amd64 go install -v -a --ldflags="-s"

ENTRYPOINT ["/go/bin/unico"]
