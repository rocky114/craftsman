CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `tel` char(11) NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱',
  `balance` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '余额',
  `points` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '1: 冻结, 2: 删除',
  `original_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '来源用户ID',
  `is_admin` tinyint(255) unsigned NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tel` (`tel`),
  KEY `original_id` (`original_id`),
  KEY `email` (`email`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

CREATE TABLE `university` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '学校名称',
  `website_address` char(32) DEFAULT NULL COMMENT '学校网址',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学校基础信息表';