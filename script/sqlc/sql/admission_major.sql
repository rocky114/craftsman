-- name: CreateAdmissionMajor :exec
INSERT INTO admission_major (
    major, province, subject_type, admission_time, duration, max_score, min_score, average_score
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);