package users

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users(username, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO NOTHING;`

	_, err := r.cluster.Conn.Exec(ctx, query, user.Username, user.ChatID)
	if err != nil {
		return fmt.Errorf("userRepo.Create: %w", err)
	}

	return nil
}
