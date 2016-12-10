FROM golang:1.7

RUN mkdir -p /go/src/github.com/alexandrevicenzi/go-chat
WORKDIR /go/src/github.com/alexandrevicenzi/go-chat

RUN go get golang.org/x/net/websocket
