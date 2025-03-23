package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) GetAllSignals() ([]domain.Signal, error) {
	query := "SELECT id, address, port FROM signals;"

	rows, err := r.cluster.Conn.Query(context.Background(), query)
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
