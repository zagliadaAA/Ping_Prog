# Используем официальный образ Golang
FROM golang:1.23-alpine AS builder

# Устанавливаем необходимые зависимости
RUN apk add --no-cache make git

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код проекта
COPY . .

# Устанавливаем goose для миграций
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Финальный образ
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk add --no-cache postgresql-client

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранное приложение
COPY --from=builder /app/main .
COPY --from=builder /app/db/migrations ./db/migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/cmd/.env .env

# Выдаем права на запуск
RUN chmod +x main

# Команда запуска
CMD sh -c "goose -dir ./db/migrations postgres \"$DATABASE_URL\" up && ./main"