# Этап 1: билд
FROM golang:1.24.2-alpine AS builder
WORKDIR /app

# Копируем модули из корня
COPY go.mod ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник сервиса
RUN go build -o ride-service ./cmd/ride-service

# Этап 2: минимальный образ для запуска
FROM alpine:latest
WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/ride-service .

# Копируем конфиг файл
COPY internal/config/config.yaml /app/config/config.yaml

# Указываем переменную окружения
ENV CONFIG_PATH=/app/config/config.yaml

CMD ["./ride-service"]