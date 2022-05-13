-- DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` INT unsigned PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键id',
    `username` varchar(20) NOT NULL DEFAULT 'Mike' COMMENT '用户名',
    `nickname` varchar(20) DEFAULT NULL COMMENT '昵称',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `age` TINYINT DEFAULT NULL COMMENT '年龄',
    `gender` TINYINT DeFAULT NULL COMMENT '性别',
    `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像地址',
    `phone` char(11) DEFAULT NULL COMMENT '电话',
    `state` TINYINT DEFAULT NULL COMMENT '状态',
    `ip` INT unsigned DEFAULT NULL COMMENT 'IP地址',
    `last_login` datetime(0) DEFAULT NULL COMMENT '最后登陆时间',
    `update_at` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `create_at` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `delete_at` datetime(0) DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `unique_idx_username` (`username`),
    UNIQUE KEY `unique_idx_email` (`email`),
    UNIQUE KEY `unique_idx_phone` (`phone`),
    KEY `idx_delete_at` (`delete_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;