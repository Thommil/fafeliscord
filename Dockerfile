FROM golang:1.12

WORKDIR /go
ENV GOPATH /go

RUN go get github.com/thommil/fafeliscord

CMD ["/go/bin/fafeliscord"]
