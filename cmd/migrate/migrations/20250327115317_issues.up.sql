CREATE TABLE IF NOT EXISTS issues (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `key` VARCHAR(255) NOT NULL UNIQUE,
    `summary` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `project_key` VARCHAR(255) NOT NULL,
    `reporter` VARCHAR(255) NOT NULL,
    `assignee` VARCHAR(255) NOT NULL,
    `status` ENUM('open', 'in_progress', 'resolved') NOT NULL DEFAULT 'open',
    `issueType` ENUM('bug', 'task', 'story') NOT NULL DEFAULT 'task',
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `sprint_id` INT UNSIGNED DEFAULT NULL,
    FOREIGN KEY (`project_key`) REFERENCES `projects`(`project_key`) ON DELETE CASCADE,
    FOREIGN KEY (`sprint_id`) REFERENCES `sprints`(`id`) ON DELETE SET NULL
);
