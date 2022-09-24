-- +migrate Up

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;