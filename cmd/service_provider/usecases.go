package service_provider

import (
	"ping_prog/internal/usecase/signal_usecase"
)

func (sp *ServiceProvider) GetSignalUseCase() *signal_usecase.UseCase {
	if sp.signalUseCase == nil {
		sp.signalUseCase = signal_usecase.NewUseCase(sp.GetSignalRepository())
	}

	return sp.signalUseCase
}
