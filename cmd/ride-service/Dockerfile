# Этап 1: Сборка
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник ride-service
RUN CGO_ENABLED=0 GOOS=linux go build -o ride-service ./cmd/ride-service

# Этап 2: Финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/ride-service .

# Копируем папку internal со всем содержимым
COPY --from=builder /app/internal ./internal

# Устанавливаем переменную окружения для пути к конфигу
ENV CONFIG_PATH=/app/internal/config/config.yaml

# Запускаем сервису
CMD ["./ride-service"]