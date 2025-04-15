package projectscopes

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

func (s *Store) CreateScope(scope types.Scope) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Insert scope
	res, err := tx.Exec("INSERT INTO scopes (name, description) VALUES (?, ?)", scope.Name, scope.Description)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert scope: %v", err)
	}

	scopeID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve scope id: %v", err)
	}

	// Associate projects with the scope
	for _, projectKey := range scope.Projects {
		_, err := tx.Exec("INSERT INTO project_scope (scope_id, project_key) VALUES (?, ?)", scopeID, projectKey)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to associate project %s: %v", projectKey, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *Store) AddProjectToScope(scopeID int, projectKey string) error {
	_, err := s.db.Exec(`
		INSERT INTO project_scope (scope_id, project_key)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE project_key = project_key -- prevent duplicates gracefully
	`, scopeID, projectKey)

	if err != nil {
		return fmt.Errorf("failed to add project to scope: %v", err)
	}
	return nil
}
