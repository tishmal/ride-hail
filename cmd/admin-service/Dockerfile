# Этап 1: билд
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY  go.mod ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/admin-service

RUN go build -o admin-service .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/admin-service/admin-service .

CMD ["./admin-service"]
