package results

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *ResultRepo) Create(ctx context.Context, res *domain.Result) error {
	query := "INSERT INTO ping_results(result, statistic, signal_id, user_id, created_at) VALUES ($1, $2, $3, $4, $5);"

	_, err := r.cluster.Conn.Exec(ctx, query, res.Result, res.Statistic, res.IDSignal, res.IDUser, res.CreatedAt)

	if err != nil {
		return fmt.Errorf("createResult: error creating result: %w", err)
	}

	return nil
}
