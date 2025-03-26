package result_usecase

import (
	"context"
	"fmt"

	"ping_prog/internal/domain"
)

func (uc *UseCase) DeleteResultsForSignal(ctx context.Context, signal *domain.Signal) error {
	err := uc.resultRepo.DeleteResultsForSignal(ctx, signal)
	if err != nil {
		return fmt.Errorf("resultRepo.DeleteResultForSignal: %w", err)
	}

	return nil
}
