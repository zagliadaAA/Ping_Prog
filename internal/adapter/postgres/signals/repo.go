package signals

import (
	"ping_prog/internal/adapter/postgres/config"
)

type Repo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
