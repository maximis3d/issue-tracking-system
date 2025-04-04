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

func (s *Store) EndCurrentStandUp(standup types.Standup) error {
	_, err := s.db.Exec(`
	UPDATE standups
	SET end_time = NOW()
	WHERE project_key = ? AND end_time IS NULL
	ORDER BY start_time DESC
	LIMIT 1`, standup.ProjectKey)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetActiveStandup(standup types.Standup) (*types.Standup, error) {
	row := s.db.QueryRow(`
		SELECT id, project_key, start_time, end_time, created_at 
		FROM standups 
		WHERE project_key = ? AND end_time IS NULL 
		ORDER BY start_time DESC 
		LIMIT 1
	`, standup.ProjectKey)

	err := row.Scan(&standup.ID, &standup.ProjectKey, &standup.StartTime, &standup.EndTime, &standup.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &standup, nil
}
