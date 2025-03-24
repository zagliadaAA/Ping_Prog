package signal_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetActiveSignalsGroupedByUser(ctx context.Context) (map[int][]domain.Signal, error) {
	signals, err := uc.signalRepo.GetActiveSignalsGroupedByUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("signalRepo.GetActiveSignalsGroupedByUser: %w", err)
	}

	return signals, nil
}
