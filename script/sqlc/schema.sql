CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `telephone` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '电话',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `balance` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '余额',
  `points` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1: 冻结, 2: 删除',
  `original_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '来源用户ID',
  `is_admin` tinyint(255) unsigned NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tel` (`telephone`),
  KEY `original_id` (`original_id`),
  KEY `email` (`email`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `school` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '学校名称',
  `code` varchar(20) NOT NULL DEFAULT '' COMMENT '学校标识码',
  `department` varchar(30) NOT NULL DEFAULT '' COMMENT '主管部门',
  `location` varchar(50) NOT NULL DEFAULT '' COMMENT '所在地',
  `level` varchar(20) NOT NULL DEFAULT '' COMMENT '办学层次',
  `website` varchar(200) NOT NULL DEFAULT '' COMMENT '学校网址',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学校基础信息表';

CREATE TABLE `admission_major` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `major` varchar(250) NOT NULL DEFAULT '' COMMENT '专业',
  `province` varchar(50) NOT NULL DEFAULT '' COMMENT '省份',
  `subject_type` varchar(100) NOT NULL DEFAULT '' COMMENT '科类',
  `admission_time` char(4) NOT NULL COMMENT '录取时间',
  `duration` smallint(5) unsigned NOT NULL DEFAULT '4' COMMENT '学制',
  `max_score` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '最高分',
  `min_score` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '最低分',
  `average_score` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '平均分',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='录取专业';