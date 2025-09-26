# Этап 1: билд
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/driver-service

RUN go build -o driver-service .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/driver-service/driver-service .

CMD ["./driver-service"]
