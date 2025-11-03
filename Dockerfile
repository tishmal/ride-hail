# ==============================================
# Универсальный Dockerfile
# ==============================================

FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata bash

WORKDIR /app

# Аргумент для выбора сервиса (ride-service / driver-service / admin-service)
ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

# Копируем go.mod и зависимости
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Копируем исходный код
COPY . .

# Сборка статического бинарника
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -a -installsuffix cgo \
    -ldflags="-w -s" \
    -trimpath \
    -o /${SERVICE_NAME} ./cmd/${SERVICE_NAME}

# ==============================================
# Финальный минимальный образ
# ==============================================
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata wget && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}
ENV SERVICE_PORT=8080

# Копируем бинарник с нужными правами
COPY --from=builder --chmod=755 /${SERVICE_NAME} /app/service

# Переключаемся на непривилегированного пользователя
USER appuser

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${SERVICE_PORT}/health || exit 1

EXPOSE ${SERVICE_PORT}

LABEL org.opencontainers.image.title="Ride-Hail ${SERVICE_NAME}" \
      org.opencontainers.image.description="Universal Go microservice for ride-hailing system" \
      org.opencontainers.image.version="1.0.0"

ENTRYPOINT ["/app/service"]
