-- name: CreateQueryCondition :exec
INSERT INTO admission_query_condition (
    university_name, url, year, province, admission_type
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: GetQueryConditionByID :one
SELECT * FROM admission_query_condition 
WHERE id = ? LIMIT 1;

-- name: GetQueryConditionByYearAndName :one
SELECT * FROM admission_query_condition 
WHERE year = ? AND university_name = ? LIMIT 1;

-- name: ListQueryConditions :many
SELECT * FROM admission_query_condition 
ORDER BY create_time DESC;

-- name: ListQueryConditionsByYear :many
SELECT * FROM admission_query_condition 
WHERE year = ?
ORDER BY create_time DESC;

-- name: UpdateQueryCondition :exec
UPDATE admission_query_condition
SET
    university_name = ?,
    url = ?,
    year = ?,
    province = ?,
    admission_type = ?
WHERE id = ?;

-- name: DeleteQueryCondition :exec
DELETE FROM admission_query_condition
WHERE id = ?;
