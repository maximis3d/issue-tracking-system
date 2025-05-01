package issue

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/maximis3d/issue-tracking-system/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateIssue(issue types.Issue) error {
	var issueCount int
	err := s.db.QueryRow("SELECT COUNT(*) FROM issues WHERE project_key = ?", issue.ProjectKey).Scan(&issueCount)
	if err != nil {
		return fmt.Errorf("failed to get issue count: %v", err)
	}

	issueNumber := issueCount + 1
	issueKey := fmt.Sprintf("%s-%03d", issue.ProjectKey, issueNumber)

	_, err = s.db.Exec("INSERT INTO issues (`key`, summary, description, project_key, reporter, assignee, status, issueType) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		issueKey, issue.Summary, issue.Description, issue.ProjectKey, issue.Reporter, issue.Assignee, issue.Status, issue.IssueType)
	if err != nil {
		return fmt.Errorf("failed to insert issue: %v", err)
	}

	_, err = s.db.Exec("UPDATE projects SET issue_count = issue_count + 1 WHERE project_key = ?", issue.ProjectKey)
	if err != nil {
		return fmt.Errorf("failed to increment issue count: %v", err)
	}

	return nil
}

func (s *Store) UpdateIssue(issue types.Issue) error {
	var currentStatus string
	err := s.db.QueryRow("SELECT status FROM issues WHERE id = ?", issue.ID).Scan(&currentStatus)
	if err != nil {
		return fmt.Errorf("failed to fetch current issue status: %v", err)
	}

	// Enforce WIP limit when moving to in_progress
	if issue.Status == "in_progress" && currentStatus != "in_progress" {
		var wipLimit int
		var inProgressCount int

		err = s.db.QueryRow("SELECT wip_limit FROM projects WHERE project_key = ?", issue.ProjectKey).Scan(&wipLimit)
		if err != nil {
			return fmt.Errorf("failed to fetch WIP limit: %v", err)
		}

		err = s.db.QueryRow("SELECT COUNT(*) FROM issues WHERE project_key = ? AND status = 'in_progress'", issue.ProjectKey).Scan(&inProgressCount)
		if err != nil {
			return fmt.Errorf("failed to fetch in-progress issues count: %v", err)
		}

		if inProgressCount >= wipLimit {
			return fmt.Errorf("too many issues in progress, the WIP limit is %d", wipLimit)
		}
	}

	// Prepare dynamic update for timestamps
	query := `
		UPDATE issues 
		SET summary = ?, description = ?, project_key = ?, reporter = ?, assignee = ?, status = ?, issueType = ?, updatedAt = NOW()`

	args := []interface{}{
		issue.Summary,
		issue.Description,
		issue.ProjectKey,
		issue.Reporter,
		issue.Assignee,
		issue.Status,
		issue.IssueType,
	}

	// Add started_at if moving to in_progress
	if issue.Status == "in_progress" && currentStatus != "in_progress" {
		query += `, started_at = NOW()`
	}

	// Add finished_at if moving to done
	if issue.Status == "resolved" && currentStatus != "resolved" {
		query += `, finished_at = NOW()`
	}

	query += ` WHERE id = ?`
	args = append(args, issue.ID)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update issue: %v", err)
	}

	return nil
}

func (s *Store) GetIssueByID(id int) (*types.Issue, error) {
	i := &types.Issue{}

	// Query to get the issue details, including started_at and finished_at
	err := s.db.QueryRow("SELECT id, `key`, summary, description, project_key, reporter, assignee, status, issueType, createdAt, updatedAt, started_at, finished_at FROM issues WHERE id = ?", id).Scan(
		&i.ID,
		&i.Key,
		&i.Summary,
		&i.Description,
		&i.ProjectKey,
		&i.Reporter,
		&i.Assignee,
		&i.Status,
		&i.IssueType,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.StartedAt,
		&i.FinishedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("issue with ID %d not found", id)
		}
		return nil, err
	}

	// Calculate cycle time if both started_at and finished_at are available
	if i.StartedAt.Valid && i.FinishedAt.Valid {
		duration := i.FinishedAt.Time.Sub(i.StartedAt.Time)
		i.CycleTime = duration.String()
	}

	return i, nil
}

func (s *Store) GetIssuesByProject(projectKey string) ([]types.Issue, error) {
	rows, err := s.db.Query("SELECT id, `key`, summary, description, project_key, reporter, assignee, status, issueType, updatedAt, started_at, finished_at FROM issues WHERE project_key=?", projectKey)
	if err != nil {
		return nil, fmt.Errorf("failed to query issues: %v", err)
	}
	defer rows.Close()

	var issues []types.Issue

	for rows.Next() {
		var i types.Issue
		err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Summary,
			&i.Description,
			&i.ProjectKey,
			&i.Reporter,
			&i.Assignee,
			&i.Status,
			&i.IssueType,
			&i.UpdatedAt,
			&i.StartedAt,
			&i.FinishedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan issue row: %v", err)
		}
		if i.StartedAt.Valid && i.FinishedAt.Valid {
			duration := i.FinishedAt.Time.Sub(i.StartedAt.Time)
			i.CycleTime = duration.String()
		}
		issues = append(issues, i)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return issues, nil
}

func (s *Store) GetAverageCycleTime(projectKey string) (time.Duration, error) {
	query := `
		SELECT started_at, finished_at 
		FROM issues 
		WHERE project_key = ? 
		AND started_at IS NOT NULL 
		AND finished_at IS NOT NULL
	`

	rows, err := s.db.Query(query, projectKey)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch issue cycle times: %v", err)
	}
	defer rows.Close()

	var totalDuration time.Duration
	var count int

	for rows.Next() {
		var startedAt, finishedAt time.Time
		if err := rows.Scan(&startedAt, &finishedAt); err != nil {
			return 0, fmt.Errorf("failed to scan row: %v", err)
		}
		totalDuration += finishedAt.Sub(startedAt)
		count++
	}

	if count == 0 {
		return 0, fmt.Errorf("no completed issues with valid cycle time for project %s", projectKey)
	}

	average := totalDuration / time.Duration(count)
	return average, nil
}
