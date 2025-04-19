CREATE TABLE `permissions` (
    `id` bigint  NOT NULL DEFAULT 0,
    `role_id` INT NOT NULL DEFAULT '0',
    `permission_id` INT NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    INDEX `idx_role`(`role_id`, `permission_id`) USING BTREE
) ENGINE=InnoDB;
