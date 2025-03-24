package signal_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

type CreateSignalReq struct {
	Address    string
	Port       int
	UserChatId int
}

func (uc *UseCase) Create(ctx context.Context, req CreateSignalReq) error {
	signal := domain.NewSignal(req.Address, req.Port)

	user, err := uc.userRepo.GetByChatID(ctx, req.UserChatId)
	if err != nil {
		return fmt.Errorf("userRepo.GetByChatID: %w", err)
	}

	signal.IDUser = int(user.ID)

	err = uc.signalRepo.Create(ctx, signal)
	if err != nil {
		return fmt.Errorf("signalRepo.Create: %w", err)
	}

	return nil
}
