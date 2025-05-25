-- 插入操作
-- name: CreateAdmissionScoreLine :exec
INSERT INTO admission_score_line (
    year,
    province,
    university_name,
    admission_batch,
    admission_type,
    admission_region,
    subject_category,
    major_group,
    lowest_score,
    lowest_score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- 删除操作
-- name: DeleteAdmissionScoreLine :exec
DELETE FROM admission_score_line
WHERE id = ?;

-- 通过ID查询
-- name: GetAdmissionScoreLineByID :one
SELECT * FROM admission_score_line
WHERE id = ?;