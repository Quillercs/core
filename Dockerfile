FROM golang:1.5

ADD . /go/src/github.com/quillercs/core

WORKDIR /go/src/github.com/quillercs/core
RUN go get ./...

EXPOSE 3000

CMD ["go", "run", "main.go"]
