package signal_usecase

import (
	"context"

	"ping_prog/internal/domain"
)

type UseCase struct {
	signalRepo signalRepo
	userRepo   userRepo
}

func NewUseCase(signalRepo signalRepo, userRepo userRepo) *UseCase {
	return &UseCase{
		signalRepo: signalRepo,
		userRepo:   userRepo,
	}
}

type signalRepo interface {
	Create(ctx context.Context, s *domain.Signal) error
	Delete(ctx context.Context, id int) error
	GetAllSignals(ctx context.Context, userName string) ([]domain.Signal, error)
	GetActiveSignalsGroupedByUser(ctx context.Context) (map[int][]domain.Signal, error)
}

type userRepo interface {
	GetByChatID(ctx context.Context, chatID int) (*domain.User, error)
	GetByID(ctx context.Context, userID int) (*domain.User, error)
}
