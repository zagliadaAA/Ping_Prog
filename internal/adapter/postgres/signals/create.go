package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) Create(s *domain.Signal) error {
	query := "INSERT INTO signals(address, port) VALUES ($1, $2);"

	_, err := r.cluster.Conn.Exec(context.Background(), query, s.Address, s.Port)

	if err != nil {
		return fmt.Errorf("createSignal: error creating signal: %w", err)
	}

	return nil
}
