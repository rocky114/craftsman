-- name: ListUniversities :many
SELECT * FROM university limit ? offset ?;

-- name: CountUniversities :one
SELECT COUNT(*) FROM university;

-- name: CreateUniversity :exec
INSERT INTO university (
  name, code, department, province, city, school_level, property
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateUniversityLastAdmissionTime :exec
UPDATE university SET last_admission_time = ? where code = ?