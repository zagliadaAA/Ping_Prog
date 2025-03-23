package user_usecase

import (
	"context"

	"ping_prog/internal/domain"
)

type UseCase struct {
	userRepo userRepo
}

func NewUseCase(userRepo userRepo) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

type userRepo interface {
	Create(ctx context.Context, user *domain.User) error
	GetByChatID(ctx context.Context, chatID int64) (*domain.User, error)
}
