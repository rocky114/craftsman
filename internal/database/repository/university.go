package repository

import (
	"context"
	"github.com/rocky114/craftman/internal/database/sqlc"
)

type ListUniversitiesParams struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (q *Repository) ListUniversities(ctx context.Context, arg ListUniversitiesParams) ([]sqlc.University, error) {
	query := "SELECT id, name, province, admission_website FROM university"

	args := make([]interface{}, 0, 1)
	if arg.Name != "" {
		query += " AND name = ?"
		args = append(args, arg.Name)
	}

	query += " ORDER BY id ASC LIMIT ? OFFSET ?"
	args = append(args, arg.Limit, arg.Offset)

	var items []sqlc.University
	if err := q.db.Select(&items, query, args...); err != nil {
		return nil, err
	}

	return items, nil
}

type CountUniversitiesParams struct {
	Name string `json:"name"`
}

type TotalCount struct {
	TotalCount int64 `db:"total_count"`
}

func (q *Repository) CountUniversities(ctx context.Context, arg CountUniversitiesParams) (int64, error) {
	query := "SELECT count(*) as total_count FROM university"

	args := make([]interface{}, 0, 1)
	if arg.Name != "" {
		query += " AND name = ?"
		args = append(args, arg.Name)
	}

	var result TotalCount
	if err := q.db.Get(&result, query, args...); err != nil {
		return 0, err
	}

	return result.TotalCount, nil
}
