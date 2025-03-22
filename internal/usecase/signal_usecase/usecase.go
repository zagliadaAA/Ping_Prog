package signal_usecase

import (
	"ping_prog/internal/domain"
)

type UseCase struct {
	signalRepo signalRepo
}

func NewUseCase(signalRepo signalRepo) *UseCase {
	return &UseCase{
		signalRepo: signalRepo,
	}
}

type signalRepo interface {
	Create(signal *domain.Signal) error
}
