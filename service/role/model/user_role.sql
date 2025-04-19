CREATE TABLE `user_role` (
    `id` BIGINT  NOT NULL DEFAULT 0,
    `user_id` BIGINT NOT NULL DEFAULT 0,
    `role_id` BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `role_id`)
) ENGINE=InnoDB;