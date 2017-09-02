FROM golang:1.8-wheezy
WORKDIR /go/src/gosh
ADD . /go/src/gosh
CMD go run gosh.go