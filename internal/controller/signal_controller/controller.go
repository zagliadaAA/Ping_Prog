package signal_controller

import (
	"ping_prog/internal/usecase/signal_usecase"
)

type Controller struct {
	signalUseCase signalUseCase
}

type signalUseCase interface {
	Create(req signal_usecase.CreateSignalReq) error
}

func NewController(useCase signalUseCase) *Controller {
	return &Controller{
		signalUseCase: useCase,
	}
}
