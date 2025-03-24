package result_usecase

import (
	"context"
	"fmt"
	"time"

	"ping_prog/internal/domain"
)

type CreateResultReq struct {
	Result    bool
	Statistic string
	IDSignal  int
	IDUser    int
	CreatedAt time.Time
}

func (uc *UseCase) Create(ctx context.Context, req CreateResultReq) error {
	result := domain.NewResult(req.Result, req.Statistic, req.CreatedAt)
	result.IDSignal = req.IDSignal
	result.IDUser = req.IDUser

	err := uc.resultRepo.Create(ctx, result)
	if err != nil {
		return fmt.Errorf("resultRepo.Create: %w", err)
	}

	return nil
}
