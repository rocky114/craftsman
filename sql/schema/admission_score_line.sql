CREATE TABLE `admission_score_line` (
                                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                        `year` char(4) NOT NULL DEFAULT '' COMMENT '录取年份（如2024）',
                                        `province` varchar(100) NOT NULL DEFAULT '' COMMENT '省份 江苏',
                                        `university_name` varchar(100) NOT NULL DEFAULT '' COMMENT '关联院校表',
                                        `admission_batch` varchar(100) NOT NULL DEFAULT '' COMMENT '录取批次',
                                        `admission_type` varchar(100) NOT NULL DEFAULT '' COMMENT '类型: 普通类,艺术类,国家专项,高校专项,中外合作,飞行技术,预科',
                                        `admission_region` varchar(100) NOT NULL DEFAULT '' COMMENT '定向区域',
                                        `subject_category` varchar(100) NOT NULL DEFAULT '' COMMENT '科类: 历史+不限',
                                        `major_group` varchar(100) NOT NULL DEFAULT '' COMMENT '专业组',
                                        `lowest_score` varchar(100) NOT NULL DEFAULT '' COMMENT '最低分',
                                        `lowest_score_rank` varchar(100) NOT NULL DEFAULT '' COMMENT '排名 200000名次',
                                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        PRIMARY KEY (`id`),
                                        KEY `university_year` (`university_name`,`year`)
) ENGINE=InnoDB AUTO_INCREMENT=817 DEFAULT CHARSET=utf8mb4 COMMENT='考试院公布的投档线';