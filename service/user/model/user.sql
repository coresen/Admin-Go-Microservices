CREATE TABLE `user`
(
    `id`          bigint  NOT NULL DEFAULT 0,
    `parent_id`  bigint NOT NULL DEFAULT 0,
    `username`        varchar(255)        NOT NULL DEFAULT '' COMMENT '用户姓名',
    `password`    varchar(255)        NOT NULL DEFAULT '' COMMENT '用户密码',
    `status`  tinyint(2)  NOT NULL DEFAULT 1 COMMENT '【是否正常】0:异常, 1:正常',
    `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ip ',
    `create_uid`  bigint NOT NULL DEFAULT 0,
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    `delete_time` timestamp DEFAULT NULL,
    `update_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `login_last` timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username_unique` (`username`),
    INDEX `idx_create_uid`(`create_uid`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;



