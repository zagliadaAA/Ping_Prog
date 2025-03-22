package service_provider

import (
	"ping_prog/internal/adapter/postgres/config"
	"ping_prog/internal/adapter/postgres/signals"
	"ping_prog/internal/controller/signal_controller"
	"ping_prog/internal/usecase/signal_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	signalUseCase *signal_usecase.UseCase

	signalRepo *signals.Repo

	signalController *signal_controller.Controller
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
