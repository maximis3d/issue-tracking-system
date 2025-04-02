package standups

import (
	"database/sql"

	"github.com/maximis3d/issue-tracking-system/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateStandup(standup types.Standup) error {
	_, err := s.db.Exec(`
		INSERT INTO standups (project_key, start_time, end_time, created_at) 
		VALUES (?, NOW(), NULL, NOW())`,
		standup.ProjectKey,
	)
	if err != nil {
		return err
	}
	return nil
}
