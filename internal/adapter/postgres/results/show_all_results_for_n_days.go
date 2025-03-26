package results

import (
	"context"
	"fmt"

	"ping_prog/internal/usecase/result_usecase"
)

func (r *ResultRepo) ShowAllResultsForNDays(ctx context.Context, userID int, days int) ([]*result_usecase.ResultView, error) {
	query := `
	SELECT pr.result, pr.statistic, pr.created_at, s.address, s.port
FROM ping_results pr
JOIN signals s ON pr.signal_id = s.id
WHERE pr.user_id = $1
  AND pr.created_at >= NOW() - make_interval(days => $2)
ORDER BY pr.created_at ASC;
	`

	rows, err := r.cluster.Conn.Query(ctx, query, userID, days)
	if err != nil {
		return nil, fmt.Errorf("ShowAllResultsForNDays: ошибка при запросе: %w", err)
	}
	defer rows.Close()

	var results []*result_usecase.ResultView
	for rows.Next() {
		var res result_usecase.ResultView
		err := rows.Scan(&res.Result, &res.Statistic, &res.CreatedAt, &res.Address, &res.Port)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сканировании: %w", err)
		}
		results = append(results, &res)
	}

	return results, nil
}
