CREATE TABLE `eb_user`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT 'id自增',
    `username`    varchar(32)  NOT NULL  DEFAULT '' COMMENT '用户名',
    `avatar`      varchar(255) DEFAULT NULL DEFAULT '' COMMENT '头像',
    `password`    varchar(255) NOT NULL DEFAULT '' COMMENT '密码密文',
    `mobile`      varchar(64)  DEFAULT NULL COMMENT '手机号',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
    `delete_time` datetime     DEFAULT NULL COMMENT '删除标记',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;