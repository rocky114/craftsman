-- name: ListUniversities :many
SELECT * FROM university limit ? offset ?;

-- name: CountUniversities :one
SELECT COUNT(*) FROM university;

-- name: GetUniversityLastAdmissionTime :one
SELECT last_admission_time FROM university WHERE code = ?;

-- name: CreateUniversity :exec
INSERT INTO university (
  name, code, department, province, city, school_level, property
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateUniversityLastAdmissionTime :exec
UPDATE university SET last_admission_time = ? WHERE code = ?;