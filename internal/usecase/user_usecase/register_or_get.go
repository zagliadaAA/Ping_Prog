package user_usecase

import (
	"context"

	"ping_prog/internal/domain"
)

func (uc *UseCase) RegisterOrGet(ctx context.Context, username string, chatID int64) (*domain.User, error) {
	user := domain.NewUser(username, chatID)

	_ = uc.userRepo.Create(ctx, user) // игнорируем ошибку, если уже существует

	return uc.userRepo.GetByChatID(ctx, chatID)
}
