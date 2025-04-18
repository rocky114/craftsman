CREATE TABLE `admission_query_condition` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
 `university_name` varchar(100) NOT NULL COMMENT '学校名称',
 `url` varchar(500) NOT NULL COMMENT '查询URL',
 `year` varchar(10) NOT NULL COMMENT '查询年份',
 `province` varchar(50) NOT NULL COMMENT '省份',
 `admission_type` varchar(50) NOT NULL COMMENT '录取类型(统招计划/专项计划等)',
 `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `university_name` (`university_name`,`year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='高校录取查询条件表';