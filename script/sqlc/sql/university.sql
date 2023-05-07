-- name: ListUniversities :many
SELECT * FROM university limit ? offset ?;

-- name: CountUniversities :one
SELECT COUNT(*) FROM university;

-- name: CreateUniversity :exec
INSERT INTO university (
  name, code, department, location, level, property
) VALUES (
  ?, ?, ?, ?, ?, ?
);