package signal_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetAllSignals(ctx context.Context, userName string) ([]domain.Signal, error) {
	signals, err := uc.signalRepo.GetAllSignals(ctx, userName)
	if err != nil {
		return nil, fmt.Errorf("signalRepo.GetAllSignals: %w", err)
	}

	return signals, nil
}
