package results

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *ResultRepo) DeleteResultsForSignal(ctx context.Context, signal *domain.Signal) error {
	query := `
		DELETE FROM ping_results
		WHERE signal_id = $1 AND user_id = $2;
	`

	_, err := r.cluster.Conn.Exec(ctx, query, signal.ID, signal.IDUser)
	if err != nil {
		return fmt.Errorf("DeleteResultsForSignal: ошибка при удалении результатов для сигнала: %w", err)
	}

	return nil
}
