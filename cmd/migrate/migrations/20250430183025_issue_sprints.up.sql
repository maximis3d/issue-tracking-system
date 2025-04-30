CREATE TABLE IF NOT EXISTS issue_sprints (
    issue_id INT UNSIGNED NOT NULL,
    sprint_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (issue_id, sprint_id),
    FOREIGN KEY (issue_id) REFERENCES issues(id) ON DELETE CASCADE,
    FOREIGN KEY (sprint_id) REFERENCES sprints(id) ON DELETE CASCADE
);
