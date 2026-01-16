# build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# собираем приложение
RUN go build -o myapp ./cmd/main.go

# ставим goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# runtime stage
FROM debian:bookworm-slim

WORKDIR /root/

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# копируем бинарники
COPY --from=builder /app/myapp .
COPY --from=builder /go/bin/goose /usr/local/bin/goose

COPY db/migrations ./db/migrations

CMD ["./myapp"]