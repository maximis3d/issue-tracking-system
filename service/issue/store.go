package issue

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

func (s *Store) CreateIssue(issue types.Issue) error {
	var issueCount int
	err := s.db.QueryRow("SELECT COUNT(*) FROM issues WHERE project = ?", issue.Project).Scan(&issueCount)
	if err != nil {
		return err
	}

	issueNumber := issueCount + 1
	issueKey := fmt.Sprintf("%s-%03d", issue.Project, issueNumber)

	_, err = s.db.Exec("INSERT INTO issues (`key`, summary, description, project_key, reporter, assignee, status, issueType) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		issueKey, issue.Summary, issue.Description, issue.Project, issue.Reporter, issue.Assignee, issue.Status, issue.IssueType)
	if err != nil {
		return err
	}

	return nil
}
