package repository

import (
	"context"
	"github.com/rocky114/craftman/internal/database/sqlc"
)

func (q *Repository) ListUniversities(ctx context.Context) ([]sqlc.University, error) {
	var listUniversities = `-- name: ListUniversities :many
SELECT id, name, province, admission_website, create_time, update_time FROM university
ORDER BY name
`

	var items []sqlc.University
	if err := q.db.Select(&items, listUniversities); err != nil {
		return nil, err
	}

	return items, nil
}
