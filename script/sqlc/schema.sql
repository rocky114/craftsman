CREATE TABLE `university` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '学校名称',
  `code` varchar(20) NOT NULL DEFAULT '' COMMENT '学校标识码',
  `department` varchar(30) NOT NULL DEFAULT '' COMMENT '主管部门',
  `province` varchar(50) NOT NULL COMMENT '省份',
  `city` varchar(50) NOT NULL DEFAULT '' COMMENT '所在地',
  `school_level` varchar(20) NOT NULL DEFAULT '' COMMENT '办学层次',
  `website` varchar(200) NOT NULL COMMENT '历史录取分数地址',
  `property` varchar(100) NOT NULL DEFAULT '' COMMENT '办学性质【公办，民办】',
  `last_admission_time` varchar(100) NOT NULL COMMENT '最后一次招生时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `university_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学校基础信息表';

CREATE TABLE `admission_major` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `university` varchar(100) NOT NULL DEFAULT '' COMMENT '大学',
  `college` varchar(100) NOT NULL DEFAULT '' COMMENT '学院',
  `major` varchar(250) NOT NULL DEFAULT '' COMMENT '专业',
  `select_exam` varchar(100) NOT NULL DEFAULT '' COMMENT '选考',
  `province` varchar(50) NOT NULL DEFAULT '' COMMENT '省份',
  `admission_type` varchar(100) NOT NULL DEFAULT '' COMMENT '招生类型',
  `admission_time` char(4) NOT NULL DEFAULT '' COMMENT '招生年份',
  `admission_number` varchar(100) NOT NULL DEFAULT '' COMMENT '招生人数',
  `duration` varchar(100) NOT NULL DEFAULT '' COMMENT '学制',
  `max_score` varchar(100) NOT NULL DEFAULT '' COMMENT '最高分',
  `min_score` varchar(100) NOT NULL DEFAULT '' COMMENT '最低分',
  `average_score` varchar(100) NOT NULL DEFAULT '' COMMENT '平均分',
  `province_control_score_line` varchar(100) NOT NULL DEFAULT '' COMMENT '省控制分数线',
  `score_rank` varchar(100) NOT NULL DEFAULT '' COMMENT '分数排名',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_identify` (`university`,`admission_time`,`admission_type`,`select_exam`,`major`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='录取专业';