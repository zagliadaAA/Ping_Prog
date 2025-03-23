package service_provider

import (
	"ping_prog/internal/usecase/signal_usecase"
	"ping_prog/internal/usecase/user_usecase"
)

func (sp *ServiceProvider) GetSignalUseCase() *signal_usecase.UseCase {
	if sp.signalUseCase == nil {
		sp.signalUseCase = signal_usecase.NewUseCase(sp.GetSignalRepository())
	}

	return sp.signalUseCase
}

func (sp *ServiceProvider) GetUserUseCase() *user_usecase.UseCase {
	if sp.userUseCase == nil {
		sp.userUseCase = user_usecase.NewUseCase(sp.GetUserRepository())
	}

	return sp.userUseCase
}
