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

func (s *Store) FilterTickets(project types.Project) ([]types.Issue, error) {
	var lastEndTIme sql.NullTime

	err := s.db.QueryRow(`
	SELECT end_time
	FROM standups
	WHERE project_key = ? and end_time IS NOT NULL
	ORDER BY end_time DESC
	LIMIT 1
	`, project.ProjectKey).Scan(&lastEndTIme)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(`
		SELECT id, key, summary, reporter, assignee, status, issueType
		FROM issues
		WHERE project_key = ? and updated_at >`,
		project.ProjectKey, lastEndTIme.Time,
	)

	if err != nil {
		return nil, err
	}

	var issues []types.Issue

	for rows.Next() {
		var t types.Issue

		if err := rows.Scan(&t.ID, &t.Key, &t.Summary, &t.Reporter, &t.Assignee, &t.Status, &t.IssueType); err != nil {
			return nil, err
		}
		issues = append(issues, t)
	}

	return issues, nil

}
