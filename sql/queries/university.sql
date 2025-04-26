-- name: CreateUniversity :exec
INSERT INTO university (name, province, admission_website)
VALUES (?, ?, ?);

-- name: GetUniversity :one
SELECT * FROM university
WHERE id = ? LIMIT 1;

-- name: UpdateUniversity :exec
UPDATE university
SET name = ?, province = ?, admission_website = ?
WHERE id = ?;

-- name: DeleteUniversity :exec
DELETE FROM university
WHERE id = ?;

-- name: GetUniversityByName :one
SELECT * FROM university
WHERE name = ? LIMIT 1;
