CREATE TABLE `menu` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `parent_id` INT DEFAULT NULL,
    `menu_name` VARCHAR(50) NOT NULL,
    `code` VARCHAR(128) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `method` varchar(16) NOT NULL,
    `icon` VARCHAR(50) DEFAULT NULL,
    `sort_order` INT DEFAULT 0,
    `status`  tinyint(2)  NOT NULL DEFAULT 1 COMMENT '【是否正常】0:异常, 1:正常',
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    `delete_time` timestamp DEFAULT NULL,
    `update_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB;
