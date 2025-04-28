-- name: InsertScoreDistribution :exec
INSERT INTO score_distribution (
    year, province, subject_category, score_range, same_score_count, cumulative_count
) VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteScoreDistribution :exec
DELETE FROM score_distribution WHERE id = ?;

-- name: UpdateScoreDistribution :exec
UPDATE score_distribution
SET
    year = ?,
    province = ?,
    subject_category = ?,
    score_range = ?,
    same_score_count = ?,
    cumulative_count = ?
WHERE id = ?;

-- name: GetScoreDistributionByID :one
SELECT * FROM score_distribution WHERE id = ?;