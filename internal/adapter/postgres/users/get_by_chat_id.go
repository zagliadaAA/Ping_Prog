package users

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *UserRepo) GetByChatID(ctx context.Context, chatID int) (*domain.User, error) {
	query := `SELECT id, username, chat_id FROM users WHERE chat_id = $1`

	row := r.cluster.Conn.QueryRow(ctx, query, chatID)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.ChatID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetByChatID: %w", err)
	}

	return user, nil
}
