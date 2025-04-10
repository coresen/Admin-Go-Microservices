CREATE TABLE `user_role` (
    `id` INT AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL DEFAULT '0',
    `role_id` BIGINT NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY (`user_id`, `role_id`)
) ENGINE=InnoDB;