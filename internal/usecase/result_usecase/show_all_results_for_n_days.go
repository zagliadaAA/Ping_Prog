package result_usecase

import (
	"context"
	"fmt"
	"time"
)

type ResultView struct {
	Address   string
	Port      int
	Result    bool
	Statistic string
	CreatedAt time.Time
}

func (uc *UseCase) ShowAllResultsForNDays(ctx context.Context, userID int, days int) ([]*ResultView, error) {
	results, err := uc.resultRepo.ShowAllResultsForNDays(ctx, userID, days)
	if err != nil {
		return nil, fmt.Errorf("resultRepo.ShowAllResultsForNDays: %w", err)
	}
	return results, nil
}
