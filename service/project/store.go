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
	rows, err := s.db.Query("SELECT id, name, description from projects")

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
		&project.Name,
		&project.Description,
	)
	if err != nil {
		return types.Project{}, err
	}
	return project, nil
}

func (s *Store) GetProjectByKey(key string) (*types.Project, error) {
	project := new(types.Project)

	err := s.db.QueryRow("SELECT * FROM projects WHERE key = ?", key).
		Scan(&project.ID, &project.Key, &project.Name, &project.Description, &project.ProjectLead, &project.IssueCount, &project.CreatedAt)

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
        INSERT INTO projects (key, name, description, project_lead) VALUES (?, ?, ?, ?)`,
		project.Key, project.Name, project.Description, project.ProjectLead,
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
