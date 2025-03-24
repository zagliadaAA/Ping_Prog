package service_provider

import (
	"ping_prog/internal/adapter/postgres/config"
	"ping_prog/internal/adapter/postgres/results"
	"ping_prog/internal/adapter/postgres/signals"
	"ping_prog/internal/adapter/postgres/users"
	"ping_prog/internal/bot"
	"ping_prog/internal/usecase/result_usecase"
	"ping_prog/internal/usecase/signal_usecase"
	"ping_prog/internal/usecase/user_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	signalUseCase *signal_usecase.UseCase
	userUseCase   *user_usecase.UseCase
	resultUseCase *result_usecase.UseCase

	signalRepo *signals.Repo
	userRepo   *users.UserRepo
	resultRepo *results.ResultRepo

	telegramBot *bot.Bot
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
