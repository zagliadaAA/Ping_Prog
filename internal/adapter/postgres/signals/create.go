package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) Create(s *domain.Signal) error {
	query := "INSERT INTO signals(address, port, created_at) VALUES ($1, $2, $3);"

	err := r.cluster.Conn.QueryRow(context.Background(), query, s.Address, s.Port, s.CreatedAt)

	if err != nil {
		return fmt.Errorf("createSignal: error creating signal: %w", err)
	}

	return nil
}
