-- name: ListUsers :many
SELECT * FROM user;

-- name: CreateUser :execresult
INSERT INTO user (
  username, password, email
) VALUES (
  ?, ?, ?
);