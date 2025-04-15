CREATE TABLE IF NOT EXISTS projects_scopes (
    `scope_id` INT UNSIGNED NOT NULL,
    `project_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`scope_id`, `project_id`),
    FOREIGN KEY (`scope_id`) REFERENCES `scopes`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`project_id`) REFERENCES `projects`(`id`) ON DELETE CASCADE
);
