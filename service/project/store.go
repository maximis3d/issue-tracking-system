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

func (s *Store) GetProjectByID(id int) (*types.Project, error) {
	rows, err := s.db.Query("SELECT * from projects where id = ?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	project := new(types.Project)

	for rows.Next() {
		project, err = scanRowsIntoProject(rows)
		if err != nil {
			return nil, err
		}
	}

	if project.ID == 0 {
		return nil, fmt.Errorf("project not found")
	}
	return project, nil

}

func scanRowsIntoProject(rows *sql.Rows) (*types.Project, error) {
	project := new(types.Project)

	err := rows.Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.ProjectLead,
		&project.IssueCount,
		&project.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *Store) CreateProject(project types.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name, description, projectLead, issueCount), values (?,?,?,?)", project.Name, project.Description, project.ProjectLead, project.IssueCount)

	if err != nil {
		return err
	}
	return nil
}
