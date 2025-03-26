package signal_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetByID(ctx context.Context, signalID int, userID int) (*domain.Signal, error) {
	signal, err := uc.signalRepo.GetByID(ctx, signalID, userID)
	if err != nil {
		return nil, fmt.Errorf("signalRepo.GetByID: %w", err)
	}

	return signal, nil
}
