CREATE TABLE `admission_summary` (
                                     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                     `year` char(4) NOT NULL COMMENT '录取年份',
                                     `province` varchar(100) NOT NULL COMMENT '省份',
                                     `university_name` varchar(100) NOT NULL COMMENT '高校名称',
                                     `admission_type` varchar(100) NOT NULL COMMENT '类型',
                                     `subject_category` varchar(100) NOT NULL COMMENT '科类',
                                     `highest_score` varchar(100) NOT NULL COMMENT '全校最高分',
                                     `highest_score_rank` varchar(100) NOT NULL COMMENT '最高分位次',
                                     `lowest_score` varchar(100) NOT NULL COMMENT '全校最低分',
                                     `lowest_score_rank` varchar(100) NOT NULL COMMENT '最低分位次',
                                     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=194 DEFAULT CHARSET=utf8mb4 COMMENT='高校录取分数汇总表';