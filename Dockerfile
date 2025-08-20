# build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp ./cmd/main.go

# runtime stage
FROM debian:bookworm-slim

WORKDIR /root/

# goose нужен для миграций
RUN apt-get update && apt-get install -y wget ca-certificates \
    && wget -qO - https://go.dev/dl/go1.22.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz \
    && ln -s /usr/local/go/bin/go /usr/local/bin/go \
    && go install github.com/pressly/goose/v3/cmd/goose@latest \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/myapp .
COPY db/migrations ./db/migrations

CMD ["./myapp"]
