package signals

import (
	"context"
	"fmt"
)

func (r *Repo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM signals WHERE id = $1;"

	_, err := r.cluster.Conn.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("createSignal: error creating signal: %w", err)
	}

	return nil
}
