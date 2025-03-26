package user_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetByUserName(ctx context.Context, userName string) (*domain.User, error) {
	user, err := uc.userRepo.GetByUserName(ctx, userName)
	if err != nil {
		return nil, fmt.Errorf("userRepo.GetUserByUserName: %w", err)
	}

	return user, nil
}
