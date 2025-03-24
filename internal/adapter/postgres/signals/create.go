package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) Create(ctx context.Context, s *domain.Signal) error {
	query := "INSERT INTO signals(address, port, user_id) VALUES ($1, $2, $3);"

	_, err := r.cluster.Conn.Exec(ctx, query, s.Address, s.Port, s.IDUser)

	if err != nil {
		return fmt.Errorf("createSignal: error creating signal: %w", err)
	}

	return nil
}
