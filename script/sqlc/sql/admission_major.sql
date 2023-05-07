-- name: CreateAdmissionMajor :exec
INSERT INTO admission_major (
    university,
    college,
    major,
    select_exam,
    province,
    admission_type,
    admission_time,
    admission_number,
    duration,
    max_score,
    min_score,
    average_score,
    province_control_score_line,
    score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);