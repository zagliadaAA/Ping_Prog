package user_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetByID(ctx context.Context, userID int) (*domain.User, error) {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetByID: %w", err)
	}

	return user, nil
}
