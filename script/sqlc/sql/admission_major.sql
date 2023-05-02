-- name: CreateAdmissionMajor :exec
INSERT INTO admission_major (
    college, major, select_exam, province, subject_type, admission_time, duration, max_score, min_score, average_score
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);