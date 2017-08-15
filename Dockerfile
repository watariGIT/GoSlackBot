FROM golang:1.7

MAINTAINER soma

RUN go get -u github.com/nlopes/slack
RUN mkdir /go/workspace
ADD *.go /go/workspace/
ENTRYPOINT go run /go/workspace/*.go