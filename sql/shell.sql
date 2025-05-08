-- 统计缺失历史录取数据的学校
select u.* from university u left join admission_score s on u.name = s.university_name where s.id is null

-- 更新admission_score最低分排名
UPDATE admission_score s
JOIN score_distribution d ON s.lowest_score = d.score_range AND s.subject_category = d.subject_category
SET s.lowest_score_rank = d.cumulative_count
WHERE s.admission_type NOT IN ('体育类', '艺术类');

-- 更新admission_score最高分排名
UPDATE admission_score s
JOIN score_distribution d ON s.highest_score = d.score_range AND s.subject_category = d.subject_category
SET s.highest_score_rank = d.cumulative_count
WHERE s.admission_type NOT IN ('体育类', '艺术类');


-- 同步admission_summary
INSERT INTO admission_summary (year, province, university_name, admission_type, subject_category, lowest_score, highest_score)
SELECT
    year,
    province,
    university_name,
    admission_type,
    subject_category,
    CASE
    WHEN MIN(
    CASE
    WHEN lowest_score = '' OR lowest_score IS NULL THEN 0
    ELSE CAST(lowest_score AS SIGNED)
    END
    ) = 0 THEN ''
    ELSE CAST(MIN(
    CASE
    WHEN lowest_score = '' OR lowest_score IS NULL THEN 0
    ELSE CAST(lowest_score AS SIGNED)
    END
    ) AS CHAR)
END AS lowest_score,
    CASE
        WHEN MAX(
            CASE
                WHEN highest_score = '' OR highest_score IS NULL THEN 0
                ELSE CAST(highest_score AS SIGNED)
            END
        ) = 0 THEN ''
        ELSE CAST(MAX(
            CASE
                WHEN highest_score = '' OR highest_score IS NULL THEN 0
                ELSE CAST(highest_score AS SIGNED)
            END
        ) AS CHAR)
END AS highest_score
FROM admission_score
GROUP BY year, province, university_name, admission_type, subject_category;

-- 更新admission_summary最低分排名
UPDATE admission_summary s
JOIN score_distribution d ON s.lowest_score = d.score_range AND s.subject_category = d.subject_category
SET s.lowest_score_rank = d.cumulative_count
WHERE s.admission_type NOT IN ('体育类', '艺术类');

-- 更新admission_summary最高分排名
UPDATE admission_summary s
JOIN score_distribution d ON s.highest_score = d.score_range AND s.subject_category = d.subject_category
SET s.highest_score_rank = d.cumulative_count
WHERE s.admission_type NOT IN ('体育类', '艺术类');


-- 查询缺失分数排名的数据
select * from admission_summary
where admission_type NOT IN ('体育类', '艺术类') and ((highest_score != '' and highest_score_rank = '') or (lowest_score != '' and lowest_score_rank = ''))


