package service_provider

import (
	"context"
	"log"

	"ping_prog/internal/adapter/postgres/config"
)

func (sp *ServiceProvider) GetDbCluster(ctx context.Context) *config.Cluster {
	if sp.dbCluster == nil {
		dbCluster, err := config.NewCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}

		sp.dbCluster = dbCluster
	}

	return sp.dbCluster
}
