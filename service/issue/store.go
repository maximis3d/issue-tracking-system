package issue

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

func (s *Store) CreateIssue(issue types.Issue) error {
	_, err := s.db.Exec("INSERT INTO issues (summary, description, project,reporter, assignee, issueType) VALUE (?,?,?,?,?,?)", issue.Summary, issue.Description, issue.Project, issue.Reporter, issue.Assignee, issue.IssueType)

	if err != nil {
		return err
	}
	return nil
}
