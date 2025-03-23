package signal_usecase

import (
	"fmt"

	"ping_prog/internal/domain"
)

type CreateSignalReq struct {
	Address string
	Port    int
}

func (uc *UseCase) Create(req CreateSignalReq) error {

	signal := domain.NewSignal(req.Address, req.Port)

	err := uc.signalRepo.Create(signal)
	if err != nil {
		return fmt.Errorf("signalRepo.Create: %w", err)
	}

	return nil
}
