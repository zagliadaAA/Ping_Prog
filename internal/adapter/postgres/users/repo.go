package users

import (
	"ping_prog/internal/adapter/postgres/config"
)

type UserRepo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *UserRepo {
	return &UserRepo{
		cluster: cluster,
	}
}
