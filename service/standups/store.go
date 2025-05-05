package standups

import (
	"database/sql"
	"errors"

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
	// Filters tickets from stand up, checks if a last stand up time exists and if not returns all issues.
	var lastEndTime sql.NullTime

	err := s.db.QueryRow(`
        SELECT end_time
        FROM standups
        WHERE project_key = ? and end_time IS NOT NULL
        ORDER BY end_time DESC
        LIMIT 1
    `, project.ProjectKey).Scan(&lastEndTime)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if lastEndTime.Valid {
		rows, err = s.db.Query("SELECT id, `key`, summary, reporter, assignee, status, issueType FROM issues WHERE project_key = ? AND updatedAt > ?", project.ProjectKey, lastEndTime.Time)
	} else {
		rows, err = s.db.Query("SELECT id,`key` summary, reporter, assignee, status, issueType FROM issues WHERE project_key = ?", project.ProjectKey)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// Get the last standup's end_time for a given project.

func (s *Store) GetLastStandupEndTime(projectKey string) (sql.NullTime, error) {
	var lastEndTime sql.NullTime
	err := s.db.QueryRow(`
		SELECT end_time 
		FROM standups 
		WHERE project_key = ? 
		  AND end_time IS NOT NULL 
		ORDER BY end_time DESC 
		LIMIT 1
	`, projectKey).Scan(&lastEndTime)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No last standup found, return a valid sql.NullTime with Valid=false
			return sql.NullTime{Valid: false}, nil
		}
		// Other real error
		return lastEndTime, err
	}

	return lastEndTime, nil
}

// Filter issues that were updated after the provided last standup's end_time
func (s *Store) FilterTicketsByEndTime(projectKey string, lastEndTime sql.NullTime) ([]types.Issue, error) {
	var rows *sql.Rows
	var err error

	if lastEndTime.Valid {
		rows, err = s.db.Query("SELECT id, `key`, summary, reporter, assignee, status, issueType FROM issues WHERE project_key = ? AND updatedAt > ?", projectKey, lastEndTime.Time)
	} else {
		// If no standup has ended, return all issues for the project
		rows, err = s.db.Query("SELECT id, `key`, summary, reporter, assignee, status, issueType FROM issues WHERE project_key = ?", projectKey)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []types.Issue
	for rows.Next() {
		var issue types.Issue
		if err := rows.Scan(&issue.ID, &issue.Key, &issue.Summary, &issue.Reporter, &issue.Assignee, &issue.Status, &issue.IssueType); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}

	return issues, nil
}
