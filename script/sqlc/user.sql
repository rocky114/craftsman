-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM user;

-- name: CreateUser :execresult
INSERT INTO user (
  username, tel
) VALUES (
  ?, ?
);

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;