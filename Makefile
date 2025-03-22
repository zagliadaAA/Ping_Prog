#migrate-up:
#	goose -dir ./db/migrations postgres "host=localhost user=postgres dbname=medical_center sslmode=disable"

migrate-up:
	goose -dir ./db/migrations postgres "host=127.0.0.1 port=5433 user=postgres password=123 dbname=ping_prog sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

lint:
	golangci-lint run ./...