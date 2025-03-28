package service_provider

import (
	"ping_prog/internal/adapter/postgres/results"
	"ping_prog/internal/adapter/postgres/signals"
	"ping_prog/internal/adapter/postgres/users"
)

func (sp *ServiceProvider) GetSignalRepository() *signals.Repo {
	if sp.signalRepo == nil {
		sp.signalRepo = signals.NewRepo(sp.GetDbCluster())
	}

	return sp.signalRepo
}

func (sp *ServiceProvider) GetUserRepository() *users.UserRepo {
	if sp.userRepo == nil {
		sp.userRepo = users.NewRepo(sp.GetDbCluster())
	}

	return sp.userRepo
}

func (sp *ServiceProvider) GetResultsRepository() *results.ResultRepo {
	if sp.resultRepo == nil {
		sp.resultRepo = results.NewRepo(sp.GetDbCluster())
	}

	return sp.resultRepo
}
