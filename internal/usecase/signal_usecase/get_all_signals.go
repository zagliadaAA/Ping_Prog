package signal_usecase

import (
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) GetAllSignals() ([]domain.Signal, error) {
	signals, err := uc.signalRepo.GetAllSignals()
	if err != nil {
		return nil, fmt.Errorf("signalRepo.GetAllSignals: %w", err)
	}

	return signals, nil
}
