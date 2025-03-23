package service_provider

import (
	"context"
	"log"

	"ping_prog/internal/adapter/postgres/config"
)

func (sp *ServiceProvider) GetDbCluster() *config.Cluster {
	if sp.dbCluster == nil {
		dbCluster, err := config.NewCluster(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		sp.dbCluster = dbCluster
	}

	return sp.dbCluster
}
