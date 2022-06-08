-- DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `username` VARCHAR(20) NOT NULL DEFAULT 'Mike' COMMENT '用户名',
    `nickname` VARCHAR(20) DEFAULT NULL COMMENT '昵称',
    `password` VARCHAR(500) NOT NULL COMMENT '密码',
    `age` TINYINT DEFAULT NULL COMMENT '年龄',
    `gender` TINYINT DEFAULT NULL COMMENT '性别',
    `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像地址',
    `phone` CHAR(11) DEFAULT NULL COMMENT '电话',
    `state` TINYINT DEFAULT NULL COMMENT '状态',
    `ip` INT unsigned DEFAULT NULL COMMENT 'IP地址',
    `last_login` DATETIME(0) DEFAULT NULL COMMENT '最后登陆时间',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `unique_idx_username` (`username`),
    UNIQUE KEY `unique_idx_email` (`email`),
    UNIQUE KEY `unique_idx_phone` (`phone`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `directory` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `directory_id` BIGINT unsigned NOT NULL COMMENT '文件唯一id',
    `directory_name` VARCHAR(30) NOT NULL COMMENT '文件夹名',
    `directory_type` VARCHAR(30) NOT NULL COMMENT '文件夹类型',
    `directory_level` TINYINT unsigned DEFAULT NULL COMMENT '层级',
    `directory_index` TINYINT unsigned DEFAULT NULL COMMENT '同级排序',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    KEY `directory_id` (`directory_id`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `directory_relation` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `ancestor` BIGINT NOT NULL COMMENT '起始',
    `descendant` BIGINT NOT NULL COMMENT '终止',
    `distance` TINYINT NOT NULL COMMENT '距离',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `unique_idx_anc_desc` (`ancestor`, `descendant`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `file` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `file_id` BIGINT unsigned NOT NULL COMMENT '文件唯一id',
    `file_name` VARCHAR(30) NOT NULL COMMENT '文件名',
    `file_size` INT unsigned  COMMENT '文件大小',
    `file_url` VARCHAR(100) DEFAULT NULL COMMENT '地址',
    `directory_id` BIGINT unsigned NOT NULL COMMENT '所属文件夹',
    `uploader` INT unsigned COMMENT '上传者',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    KEY `idx_file_id` (`file_id`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS `role` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `role_id` BIGINT unsigned NOT NULL COMMENT '角色唯一id',
    `role_name` VARCHAR(30) NOT NULL COMMENT '角色名',
    `role_level` TINYINT unsigned DEFAULT NULL COMMENT '层级',
    `role_index` TINYINT unsigned DEFAULT NULL COMMENT '同级排序',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    KEY `role_id` (`role_id`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `role_relation` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `ancestor` BIGINT NOT NULL COMMENT '起始',
    `descendant` BIGINT NOT NULL COMMENT '终止',
    `distance` TINYINT NOT NULL COMMENT '距离',
    `update_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` DATETIME(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` DATETIME(0) DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `unique_idx_anc_desc` (`ancestor`, `descendant`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;