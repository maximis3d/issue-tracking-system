package issue

import (
	"database/sql"
	"errors"
	"fmt"

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
		return err
	}

	issueNumber := issueCount + 1
	issueKey := fmt.Sprintf("%s-%03d", issue.ProjectKey, issueNumber)

	_, err = s.db.Exec("INSERT INTO issues (`key`, summary, description, project_key, reporter, assignee, status, issueType) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		issueKey, issue.Summary, issue.Description, issue.ProjectKey, issue.Reporter, issue.Assignee, issue.Status, issue.IssueType)
	if err != nil {
		return err
	}

	return nil
}
func (s *Store) UpdateIssue(issue types.Issue) error {
	_, err := s.db.Exec(`
		UPDATE issues 
		SET summary = ?, description = ?, project_key = ?, reporter = ?, assignee = ?, status = ?, issueType = ?, updatedAt = NOW()
		WHERE id = ?`,
		issue.Summary, issue.Description, issue.ProjectKey, issue.Reporter, issue.Assignee, issue.Status, issue.IssueType, issue.ID)

	if err != nil {
		return fmt.Errorf("failed to update issue: %v", err)
	}
	return nil
}

func (s *Store) GetIssueByID(id int) (*types.Issue, error) {
	i := &types.Issue{}

	err := s.db.QueryRow("SELECT id, `key`, summary, description, project_key, reporter, assignee, status, issueType, updatedAt FROM issues WHERE id=?", id).
		Scan(
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
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("issue with ID %d not found", id)
		}
		return nil, err
	}

	return i, nil
}
