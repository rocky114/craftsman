-- name: ListSchools :many
SELECT * FROM school limit ? offset ?;

-- name: CountSchools :one
SELECT COUNT(*) FROM school;

-- name: CreateSchool :execresult
INSERT INTO school (
  name, code, department, location, level, remark
) VALUES (
  ?, ?, ?, ?, ?, ?
);