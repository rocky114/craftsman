-- name: ListUser :many
SELECT * FROM user;

-- name: GetUser :one
SELECT id, username from user where username = ? and password = ?;

-- name: CreateUser :execresult
INSERT INTO user (
  username, password, email, telephone
) VALUES (
  ?, ?, ?, ?
);