CREATE TABLE `admission_score` (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
   `year` char(4) NOT NULL DEFAULT '' COMMENT '录取年份（如2024）',
   `university_name` varchar(100) NOT NULL DEFAULT '' COMMENT '关联院校表',
   `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份 江苏',
   `admission_type` varchar(100) NOT NULL DEFAULT '' COMMENT '类型: 普通类/高校专项/地方专项',
   `academic_category` varchar(100) NOT NULL DEFAULT '' COMMENT '科类: 历史+不限',
   `major_name` varchar(100) NOT NULL DEFAULT '' COMMENT '专业名称 计算机',
   `enrollment_quota` varchar(100) NOT NULL DEFAULT '' COMMENT '招生名额',
   `min_admission_score` varchar(100) NOT NULL DEFAULT '' COMMENT '投档分 600',
   `highest_score` varchar(100) NOT NULL DEFAULT '' COMMENT '最高分',
   `highest_score_rank` varchar(100) NOT NULL DEFAULT '' COMMENT '排名 200000名次',
   `lowest_score` varchar(100) NOT NULL DEFAULT '' COMMENT '最低分',
   `lowest_score_rank` varchar(100) NOT NULL DEFAULT '' COMMENT '排名 200000名次',
   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   KEY `university_year` (`university_name`,`year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='大学历史录取分数线';