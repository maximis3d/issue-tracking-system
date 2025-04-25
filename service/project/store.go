package project

import (
	"database/sql"
	"fmt"

	"github.com/maximis3d/issue-tracking-system/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProjects() ([]types.Project, error) {
	rows, err := s.db.Query("SELECT id, project_key, name, description, project_lead, issue_count from projects")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects = []types.Project{}

	for rows.Next() {
		project, err := scanRowsIntoProjects(rows)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)

	}
	if len(projects) == 0 {
		return nil, fmt.Errorf("no projects found")
	}
	return projects, nil

}

func scanRowsIntoProjects(rows *sql.Rows) (types.Project, error) {
	var project types.Project

	err := rows.Scan(
		&project.ID,
		&project.ProjectKey,
		&project.Name,
		&project.Description,
		&project.ProjectLead,
		&project.IssueCount,
	)
	if err != nil {
		return types.Project{}, err
	}
	return project, nil
}

func (s *Store) GetProjectByKey(key string) (*types.Project, error) {
	project := new(types.Project)

	err := s.db.QueryRow(`
        SELECT id, project_key, name, description, project_lead, issue_count, created_at
        FROM projects
        WHERE project_key = ?`, key).
		Scan(
			&project.ID,
			&project.ProjectKey,
			&project.Name,
			&project.Description,
			&project.ProjectLead,
			&project.IssueCount,
			&project.CreatedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, err
	}

	return project, nil
}

func (s *Store) CreateProject(project types.Project) error {
	if !projectLeadExists(s.db, project.ProjectLead) {
		return fmt.Errorf("project lead with id %d does not exist", project.ProjectLead)
	}

	_, err := s.db.Exec(`
        INSERT INTO projects (project_key, name, description, project_lead, wip_limit) VALUES (?, ?, ?, ?, ?)`,
		project.ProjectKey, project.Name, project.Description, project.ProjectLead, project.WIPLimit,
	)
	if err != nil {
		return fmt.Errorf("error inserting project: %w", err)
	}
	return nil
}

func projectLeadExists(db *sql.DB, userID int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userID).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
