CREATE TABLE `permissions` (
    `id` BIGINT AUTO_INCREMENT,
    `role_id` INT NOT NULL DEFAULT '0',
    `permission_id` INT NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_role`(`role_id`, `permission_id`) USING BTREE
) ENGINE=InnoDB;
