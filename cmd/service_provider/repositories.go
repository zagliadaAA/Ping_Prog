package service_provider

import (
	"context"

	"ping_prog/internal/adapter/postgres/signals"
)

func (sp *ServiceProvider) GetSignalRepository() *signals.Repo {
	if sp.signalRepo == nil {
		sp.signalRepo = signals.NewRepo(sp.GetDbCluster(context.Background()))
	}

	return sp.signalRepo
}
