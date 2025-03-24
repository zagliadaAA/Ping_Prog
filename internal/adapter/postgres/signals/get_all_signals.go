package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) GetAllSignals(ctx context.Context, userName string) ([]domain.Signal, error) {
	query := `
		SELECT s.id, s.address, s.port
		FROM signals s
		JOIN users u ON s.user_id = u.id
		WHERE u.username = $1;
	`

	rows, err := r.cluster.Conn.Query(ctx, query, userName)
	if err != nil {
		return nil, fmt.Errorf("GetAllSignals: query error: %w", err)
	}
	defer rows.Close()

	var signals []domain.Signal
	for rows.Next() {
		var s domain.Signal
		if err := rows.Scan(&s.ID, &s.Address, &s.Port); err != nil {
			return nil, fmt.Errorf("GetAllSignals: scan error: %w", err)
		}
		signals = append(signals, s)
	}

	return signals, nil
}
