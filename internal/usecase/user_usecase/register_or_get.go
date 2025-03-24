package user_usecase

import (
	"context"

	"ping_prog/internal/domain"
)

func (uc *UseCase) RegisterOrGet(ctx context.Context, username string, chatID int) (*domain.User, error) {
	user := domain.NewUser(username, chatID)

	_ = uc.userRepo.Create(ctx, user) // ignore error if already exists

	return uc.userRepo.GetByChatID(ctx, chatID)
}
