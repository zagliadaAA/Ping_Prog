package result_usecase

import (
	"context"

	"ping_prog/internal/domain"
)

type UseCase struct {
	resultRepo resultRepo
}

func NewUseCase(resultRepo resultRepo) *UseCase {
	return &UseCase{
		resultRepo: resultRepo,
	}
}

type resultRepo interface {
	Create(ctx context.Context, res *domain.Result) error
}
