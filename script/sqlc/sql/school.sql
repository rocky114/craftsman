-- name: ListSchool :many
SELECT * FROM school limit ? offset ?;

-- name: CreateSchool :execresult
INSERT INTO school (
  name, code, department, location, level, remark
) VALUES (
  ?, ?, ?, ?, ?, ?
);