CREATE TABLE `score_distribution` (
                                      `id` int(11) NOT NULL AUTO_INCREMENT,
                                      `year` varchar(10) NOT NULL DEFAULT '2024' COMMENT '年份',
                                      `province` varchar(10) NOT NULL DEFAULT '江苏' COMMENT '省份',
                                      `subject_category` varchar(10) NOT NULL COMMENT '物理等科目类, 历史等科目类',
                                      `score_range` varchar(20) NOT NULL COMMENT '分数段',
                                      `same_score_count` varchar(10) NOT NULL COMMENT '同分人数',
                                      `cumulative_count` varchar(10) NOT NULL COMMENT '累计人数',
                                      PRIMARY KEY (`id`),
                                      KEY `search` (`year`,`province`,`subject_category`,`score_range`)
) ENGINE=InnoDB AUTO_INCREMENT=403 DEFAULT CHARSET=utf8mb4 COMMENT='一段一分表';