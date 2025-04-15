CREATE TABLE IF NOT EXISTS project_scope (
    scope_id INT NOT NULL,
    project_key VARCHAR(255) NOT NULL,
    PRIMARY KEY (scope_id, project_key),
    FOREIGN KEY (scope_id) REFERENCES scopes(id) ON DELETE CASCADE,
    FOREIGN KEY (project_key) REFERENCES projects(project_key) ON DELETE CASCADE
);
