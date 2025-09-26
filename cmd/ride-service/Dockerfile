# Этап 1: билд
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Копируем модули из корня
COPY  go.mod ./
RUN go mod download

# Копируем только этот сервис
COPY . .

# Переходим в сервисную папку
WORKDIR /app/cmd/ride-service

# Собираем бинарник
RUN go build -o ride-service .

# Этап 2: минимальный образ для запуска
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/ride-service/ride-service .

CMD ["./ride-service"]
