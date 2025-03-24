package signal_usecase

import (
	"context"
	"fmt"
)

func (uc *UseCase) Delete(ctx context.Context, id int) error {
	err := uc.signalRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("signalRepo.Delete: %w", err)
	}

	return nil
}
