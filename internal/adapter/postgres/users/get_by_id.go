package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *UserRepo) GetByID(ctx context.Context, userID int) (*domain.User, error) {
	query := `SELECT id, username, chat_id FROM users WHERE id = $1`

	row := r.cluster.Conn.QueryRow(ctx, query, userID)

	var user domain.User

	err := row.Scan(&user.ID, &user.Username, &user.ChatID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Если пользователя с таким ID не существует
			return nil, fmt.Errorf("пользователь с ID %d не найден", userID)
		}
		// Если произошла другая ошибка при сканировании
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	// Возвращаем полученного пользователя
	return &user, nil
}
