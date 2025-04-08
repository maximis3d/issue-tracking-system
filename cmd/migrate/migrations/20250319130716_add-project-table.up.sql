CREATE TABLE IF NOT EXISTS projects (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `project_key` VARCHAR(255) NOT NULL UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `project_lead` INT UNSIGNED NULL,
    `issue_count` INT NOT NULL DEFAULT 0,
    `wip_limit` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`project_lead`) REFERENCES `users`(`id`) ON DELETE SET NULL
);