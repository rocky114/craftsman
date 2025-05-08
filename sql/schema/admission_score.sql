CREATE TABLE `admission_score` (
                                   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                   `year` char(4) NOT NULL DEFAULT '' COMMENT '录取年份（如2024）',
                                   `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份 江苏',
                                   `university_name` varchar(100) NOT NULL DEFAULT '' COMMENT '关联院校表',
                                   `admission_type` varchar(100) NOT NULL DEFAULT '' COMMENT '类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科',
                                   `subject_category` varchar(100) NOT NULL DEFAULT '' COMMENT '科类: 历史+不限',
                                   `subject_category_txt` varchar(100) NOT NULL DEFAULT '' COMMENT '科类文本',
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='高校专业录取分数';