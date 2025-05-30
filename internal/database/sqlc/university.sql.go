// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: university.sql

package sqlc

import (
	"context"
)

const createUniversity = `-- name: CreateUniversity :exec
INSERT INTO university (name, province, admission_website)
VALUES (?, ?, ?)
`

type CreateUniversityParams struct {
	Name             string `db:"name"`
	Province         string `db:"province"`
	AdmissionWebsite string `db:"admission_website"`
}

func (q *Queries) CreateUniversity(ctx context.Context, arg CreateUniversityParams) error {
	_, err := q.db.ExecContext(ctx, createUniversity, arg.Name, arg.Province, arg.AdmissionWebsite)
	return err
}

const deleteUniversity = `-- name: DeleteUniversity :exec
DELETE FROM university
WHERE id = ?
`

func (q *Queries) DeleteUniversity(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deleteUniversity, id)
	return err
}

const getUniversity = `-- name: GetUniversity :one
SELECT id, name, province, admission_website, create_time, update_time FROM university
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUniversity(ctx context.Context, id uint32) (University, error) {
	row := q.db.QueryRowContext(ctx, getUniversity, id)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Province,
		&i.AdmissionWebsite,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUniversityByName = `-- name: GetUniversityByName :one
SELECT id, name, province, admission_website, create_time, update_time FROM university
WHERE name = ? LIMIT 1
`

func (q *Queries) GetUniversityByName(ctx context.Context, name string) (University, error) {
	row := q.db.QueryRowContext(ctx, getUniversityByName, name)
	var i University
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Province,
		&i.AdmissionWebsite,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const updateUniversity = `-- name: UpdateUniversity :exec
UPDATE university
SET name = ?, province = ?, admission_website = ?
WHERE id = ?
`

type UpdateUniversityParams struct {
	Name             string `db:"name"`
	Province         string `db:"province"`
	AdmissionWebsite string `db:"admission_website"`
	ID               uint32 `db:"id"`
}

func (q *Queries) UpdateUniversity(ctx context.Context, arg UpdateUniversityParams) error {
	_, err := q.db.ExecContext(ctx, updateUniversity,
		arg.Name,
		arg.Province,
		arg.AdmissionWebsite,
		arg.ID,
	)
	return err
}
