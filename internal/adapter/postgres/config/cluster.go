package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Cluster struct {
	Conn *pgx.Conn
}

func NewCluster(ctx context.Context) (*Cluster, error) {
	//dsn := fmt.Sprintf("postgres://localhost:5433?dbname=ping_prog&user=postgres&password=123&sslmode=disable")
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://localhost:5433?dbname=ping_prog&user=postgres&password=123&sslmode=disable"
		//dsn = "postgres://192.168.0.177:5434?dbname=ping_prog&user=postgres&password=123&sslmode=disable"
	}

	fmt.Printf("dsn: %s\n", dsn)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect: %w", err)
	}

	return &Cluster{Conn: conn}, nil
}
