package user_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) Create(ctx context.Context, userName string, chatID int) (*domain.User, error) {
	user, err := uc.userRepo.GetByUserName(ctx, userName)
	if err != nil {
		if user == nil {
			err = uc.userRepo.Create(ctx, domain.NewUser(userName, chatID))
			if err != nil {
				return nil, fmt.Errorf("userRepo.Create: %w", err)
			}
		}

		return nil, fmt.Errorf("userRepo.GetUserByUserName: %w", err)
	}

	return user, nil
}
