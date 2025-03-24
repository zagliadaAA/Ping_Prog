package results

import (
	"ping_prog/internal/adapter/postgres/config"
)

type ResultRepo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *ResultRepo {
	return &ResultRepo{
		cluster: cluster,
	}
}
