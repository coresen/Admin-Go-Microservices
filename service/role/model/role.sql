CREATE TABLE `role` (
    `id` BIGINT  NOT NULL DEFAULT 0,
    `parent_id` BIGINT NOT NULL,
    `role_name` VARCHAR(128) NOT NULL DEFAULT '' UNIQUE,
    `description` VARCHAR(255) NOT NULL DEFAULT '',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `delete_at` TIMESTAMP DEFAULT NULL,
    `updated_at` TIMESTAMP DEFAULT NULL,
    `status` INT NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;
