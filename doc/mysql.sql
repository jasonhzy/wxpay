CREATE DATABASE IF NOT EXISTS beego DEFAULT charset utf8mb4 COLLATE utf8mb4_unicode_ci;

use beego;

CREATE TABLE if NOT EXISTS user (
  `id` INT not null auto_increment,
  `username`  VARCHAR(100) NOT NULL default '',
  `password`  VARCHAR(32) NOT NULL default '',
  `nickname` VARCHAR(100) NOT NULL default '',
  `age` INT DEFAULT 0,
  `sex` TINYINT DEFAULT 0 COMMENT '0 male 1 female',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE = InnoDB DEFAULT charset=utf8mb4;