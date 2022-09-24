-- name: ListUser :many
SELECT * FROM user;

-- name: GetUser :one
select id, username from user where username = ? and password = ?;

-- name: CreateUser :execresult
INSERT INTO user (
  username, password, email
) VALUES (
  ?, ?, ?
);