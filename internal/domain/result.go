package domain

import (
	"time"
)

type Result struct {
	ID        int
	Result    bool
	Statistic string
	IDSignal  int
	IDUser    int
	CreatedAt time.Time
}

func NewResult(result bool, statistic string, createdAt time.Time) *Result {
	return &Result{
		Result:    result,
		Statistic: statistic,
		CreatedAt: createdAt,
	}
}
