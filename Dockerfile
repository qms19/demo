FROM golang:1.17.0-alpine3.14 as builder
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN  export GOPROXY="https://goproxy.cn,direct" \
     go mod download

COPY main.go .
RUN  go build -o web-demo main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/web-demo /app/web-demo
ENTRYPOINT ./web-demo

EXPOSE 9000