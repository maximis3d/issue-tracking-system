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

func (s *Store) GetProjectByName(name string) (*types.Project, error) {
	project := new(types.Project)

	err := s.db.QueryRow("SELECT * FROM projects WHERE name = ?", name).
		Scan(&project.ID, &project.Name, &project.Description, &project.ProjectLead, &project.IssueCount, &project.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, err
	}

	return project, nil
}

func (s *Store) CreateProject(project types.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name, description, project_lead) VALUES (?, ?, ?)",
		project.Name, project.Description, project.ProjectLead)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateIssue(name types.Project) error {
	_, err := s.db.Exec("UPDATE projects SET issueCount = issueCount + 1 WHERE project = ?", name)
	if err != nil {
		return err
	}
	return nil

}
