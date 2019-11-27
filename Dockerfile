FROM golang:1.12

WORKDIR /go
ENV GOPATH /go
ENV MONGO_URL="mongodb://mongo:27017"

RUN go get github.com/thommil/fafeliscord

CMD ["/go/bin/fafeliscord"]
