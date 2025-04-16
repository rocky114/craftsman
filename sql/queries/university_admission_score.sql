-- name: CreateAdmissionScore :exec
INSERT INTO university_admission_score (
    university_id, year, province, admission_type, academic_category,
    major_name, enrollment_quota, min_admission_score, min_admission_rank,
    highest_score, highest_score_rank, lowest_score, lowest_score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetAdmissionScoreByID :one
SELECT * FROM university_admission_score 
WHERE id = ? LIMIT 1;

-- name: ListAdmissionScores :many
SELECT * FROM university_admission_score 
ORDER BY create_time DESC;

-- name: ListAdmissionScoresByUniversity :many
SELECT * FROM university_admission_score 
WHERE university_id = ?
ORDER BY year DESC, create_time DESC;

-- name: ListAdmissionScoresByYearAndProvince :many
SELECT * FROM university_admission_score 
WHERE year = ? AND province = ?
ORDER BY university_id, admission_type, academic_category;

-- name: ListAdmissionScoresByTypeAndCategory :many
SELECT * FROM university_admission_score 
WHERE admission_type = ? AND academic_category = ?
ORDER BY year DESC, province;

-- name: UpdateAdmissionScore :exec
UPDATE university_admission_score
SET 
    university_id = ?,
    year = ?,
    province = ?,
    admission_type = ?,
    academic_category = ?,
    major_name = ?,
    enrollment_quota = ?,
    min_admission_score = ?,
    min_admission_rank = ?,
    highest_score = ?,
    highest_score_rank = ?,
    lowest_score = ?,
    lowest_score_rank = ?
WHERE id = ?;

-- name: DeleteAdmissionScore :exec
DELETE FROM university_admission_score
WHERE id = ?;
