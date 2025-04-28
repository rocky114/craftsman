package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/rocky114/craftman/internal/database/sqlc"
)

type UniversityQueryParams struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (q *Repository) buildUniversityQuery(baseQuery string, arg UniversityQueryParams) (string, []interface{}) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	args := make([]interface{}, 0)
	conditions := make([]string, 0)

	if arg.Name != "" {
		conditions = append(conditions, "name LIKE ?")
		args = append(args, "%"+arg.Name+"%")
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	return queryBuilder.String(), args
}

func (q *Repository) ListUniversities(ctx context.Context, arg UniversityQueryParams) ([]sqlc.University, error) {
	baseQuery := "SELECT id, name, province, admission_website FROM university"
	query, args := q.buildUniversityQuery(baseQuery, arg)

	query += " ORDER BY id ASC"
	if arg.Limit > 0 {
		query += " LIMIT ?"
		args = append(args, arg.Limit)
	}
	if arg.Offset > 0 {
		query += " OFFSET ?"
		args = append(args, arg.Offset)
	}

	var items []sqlc.University
	if err := q.db.SelectContext(ctx, &items, query, args...); err != nil {
		return nil, fmt.Errorf("ListUniversities failed: %w", err)
	}
	return items, nil
}

func (q *Repository) CountUniversities(ctx context.Context, arg UniversityQueryParams) (int64, error) {
	baseQuery := "SELECT COUNT(*) AS total_count FROM university"
	query, args := q.buildUniversityQuery(baseQuery, arg)

	var total int64
	if err := q.db.GetContext(ctx, &total, query, args...); err != nil {
		return 0, fmt.Errorf("CountUniversities failed: %w", err)
	}
	return total, nil
}
