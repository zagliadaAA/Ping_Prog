package signals

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *Repo) GetActiveSignalsGroupedByUser(ctx context.Context) (map[int][]domain.Signal, error) {
	query := `
        SELECT id, address, port, user_id 
        FROM signals 
        WHERE active = TRUE;
    `

	rows, err := r.cluster.Conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("GetActiveSignalsGroupedByUser: ошибка запроса: %w", err)
	}
	defer rows.Close()

	// Мапа для группировки сигналов по user_id
	activeSignalsByUser := make(map[int][]domain.Signal)

	// Проходим по строкам результата запроса
	for rows.Next() {
		var signal domain.Signal
		if err := rows.Scan(&signal.ID, &signal.Address, &signal.Port, &signal.IDUser); err != nil {
			return nil, fmt.Errorf("GetActiveSignalsGroupedByUser: ошибка сканирования: %w", err)
		}

		// Добавляем сигнал в мапу по user_id
		activeSignalsByUser[signal.IDUser] = append(activeSignalsByUser[signal.IDUser], signal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetActiveSignalsGroupedByUser: ошибка итерации по строкам: %w", err)
	}

	return activeSignalsByUser, nil
}
