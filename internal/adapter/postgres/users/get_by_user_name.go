package users

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *UserRepo) GetByUserName(ctx context.Context, userName string) (*domain.User, error) {
	query := `SELECT id, username, chat_id FROM users WHERE username = $1`

	row := r.cluster.Conn.QueryRow(ctx, query, userName)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.ChatID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetByUserName: %w", err)
	}

	return user, nil
}
