package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) GetByID(ctx context.Context, signalID int, userID int) (*domain.Signal, error) {
	query := `
		SELECT id, address, port, user_id
		FROM signals
		WHERE id = $1 AND user_id = $2`

	row := r.cluster.Conn.QueryRow(ctx, query, signalID, userID)

	var signal domain.Signal
	err := row.Scan(&signal.ID, &signal.Address, &signal.Port, &signal.IDUser)
	if err != nil {
		return nil, fmt.Errorf("GetByID: ошибка при чтении результата: %w", err)
	}

	return &signal, nil
}
