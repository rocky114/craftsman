// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: univeristy.sql

package storage

import (
	"context"
)

const listUniversity = `-- name: ListUniversity :many
SELECT id, name, website_address, create_time, update_time FROM university
`

func (q *Queries) ListUniversity(ctx context.Context) ([]University, error) {
	rows, err := q.db.QueryContext(ctx, listUniversity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []University
	for rows.Next() {
		var i University
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.WebsiteAddress,
			&i.CreateTime,
			&i.UpdateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}