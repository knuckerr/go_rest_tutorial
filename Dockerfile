FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc software-properties-common

RUN  add-apt-repository ppa:longsleep/golang-backports
RUN  apt update -y
RUN  apt install golang-go -y

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY . /api
WORKDIR /api

RUN go mod download

RUN go build main.go

EXPOSE 8080

CMD ["./main"]
