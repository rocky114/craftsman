INSERT INTO admission_summary (year, university_name, admission_type, subject_category, lowest_score, highest_score)
SELECT
    '2024',
    university_name,
    admission_type,
    subject_category,
    MIN(
            CASE
                WHEN lowest_score = '' OR lowest_score IS NULL THEN 0
                ELSE CAST(lowest_score AS SIGNED)
                END
    ) AS lowest_score,
    NULLIF(MAX(
                   CASE
                       WHEN highest_score = '' OR highest_score IS NULL THEN 0
                       WHEN highest_score = '0' THEN 0  -- 将'0'转换为0以便NULLIF处理
                       ELSE CAST(highest_score AS SIGNED)
                       END
           ), 0) AS highest_score  -- 将结果为0的转换为NULL
FROM admission_score
WHERE admission_type NOT IN ('艺术类', '体育类')
GROUP BY university_name, admission_type, subject_category;



INSERT INTO admission_summary (year, university_name, admission_type, lowest_score, highest_score)
SELECT
    '2024',
    university_name,
    admission_type,
    MIN(
            CASE
                WHEN lowest_score = '' OR lowest_score IS NULL THEN 0
                ELSE CAST(lowest_score AS SIGNED)
                END
    ) AS lowest_score,
    NULLIF(MAX(
                   CASE
                       WHEN highest_score = '' OR highest_score IS NULL THEN 0
                       WHEN highest_score = '0' THEN 0  -- 将'0'转换为0以便NULLIF处理
                       ELSE CAST(highest_score AS SIGNED)
                       END
           ), 0) AS highest_score  -- 将结果为0的转换为NULL
FROM admission_score
WHERE admission_type IN ('艺术类', '体育类')
GROUP BY university_name, admission_type;


-- 统计缺失历史录取数据的学校
select u.* from university u left join admission_score s on u.name = s.university_name where s.id is null
