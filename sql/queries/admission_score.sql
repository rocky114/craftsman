-- name: CreateAdmissionScore :exec
INSERT INTO admission_score (
    id, university_name, year, province, admission_type, academic_category,
    major_name, enrollment_quota, min_admission_score, min_admission_rank,
    highest_score, highest_score_rank, lowest_score, lowest_score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetAdmissionScoreByID :one
SELECT * FROM admission_score 
WHERE id = ? LIMIT 1;

-- name: ListAdmissionScores :many
SELECT * FROM admission_score 
ORDER BY create_time DESC;

-- name: ListAdmissionScoresByUniversity :many
SELECT * FROM admission_score 
WHERE id = ?
ORDER BY year DESC, create_time DESC;

-- name: ListAdmissionScoresByYearAndProvince :many
SELECT * FROM admission_score 
WHERE year = ? AND province = ?
ORDER BY id, admission_type, academic_category;

-- name: ListAdmissionScoresByTypeAndCategory :many
SELECT * FROM admission_score 
WHERE admission_type = ? AND academic_category = ?
ORDER BY year DESC, province;

-- name: DeleteAdmissionScore :exec
DELETE FROM admission_score
WHERE id = ?;

-- name: DeleteAdmissionScoreByYearAndUniversity :exec
DELETE FROM admission_score
WHERE year = ? AND university_name = ?;
