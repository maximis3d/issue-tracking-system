CREATE TABLE IF NOT EXISTS project_scope (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `project_key` VARCHAR(255) NOT NULL,
    `scope_id` INT NOT NULL,
    FOREIGN KEY (`project_key`) REFERENCES `projects`(`project_key`) ON DELETE CASCADE,
    FOREIGN KEY (`scope_id`) REFERENCES `scopes`(`id`) ON DELETE CASCADE
);
