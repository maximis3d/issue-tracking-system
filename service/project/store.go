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
