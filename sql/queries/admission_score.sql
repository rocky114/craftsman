-- name: CreateAdmissionScore :exec
INSERT INTO admission_score (
    id, university_name, year, province, admission_type, subject_category,
    major_name, highest_score, highest_score_rank, lowest_score, lowest_score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetAdmissionScoreByID :one
SELECT * FROM admission_score 
WHERE id = ? LIMIT 1;

-- name: DeleteAdmissionScore :exec
DELETE FROM admission_score
WHERE id = ?;

-- name: DeleteAdmissionScoreByYearAndUniversity :exec
DELETE FROM admission_score
WHERE year = ? AND university_name = ?;
