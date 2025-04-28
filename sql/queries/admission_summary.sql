-- 插入操作
-- name: CreateAdmissionSummary :exec
INSERT INTO admission_summary (
    year,
    province,
    university_name,
    admission_type,
    subject_category,
    highest_score,
    highest_score_rank,
    lowest_score,
    lowest_score_rank
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- 删除操作
-- name: DeleteAdmissionSummary :exec
DELETE FROM admission_summary 
WHERE id = ?;

-- 通过ID查询
-- name: GetAdmissionSummaryByID :one
SELECT * FROM admission_summary 
WHERE id = ?;