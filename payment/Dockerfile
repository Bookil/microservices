FROM golang:latest AS builder

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn

WORKDIR /usr/src/app

COPY . ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payment ./cmd/main.go

CMD ["./payment"]







