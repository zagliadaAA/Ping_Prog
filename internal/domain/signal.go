package domain

import (
	"time"
)

type Signal struct {
	ID        int
	Address   string
	Port      int
	CreatedAt time.Time
}

func NewSignal(address string, port int, createdAt time.Time) *Signal {
	return &Signal{
		Address:   address,
		Port:      port,
		CreatedAt: createdAt,
	}
}
