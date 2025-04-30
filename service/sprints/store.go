package sprints

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

func (s *Store) CreateSprint(sprint types.Sprint) error {
	_, err := s.db.Exec(`
        INSERT INTO sprints (name, description, start_date, end_date, project_key)
        VALUES (?, ?, ?, ?, ?)`,
		sprint.Name, sprint.Description, sprint.StartDate, sprint.EndDate, sprint.ProjectKey,
	)
	if err != nil {
		return fmt.Errorf("error inserting sprint: %w", err)
	}
	return nil
}

func (s *Store) AddIssueToSprint(issueID, sprintID int) error {
	var sprintProjectKey string
	err := s.db.QueryRow("SELECT project_key FROM sprints WHERE id = ?", sprintID).Scan(&sprintProjectKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("sprint with ID %d not found", sprintID)
		}
		return fmt.Errorf("failed to fetch sprint project key: %v", err)
	}

	var issueProjectKey string
	err = s.db.QueryRow("SELECT project_key FROM issues WHERE id = ?", issueID).Scan(&issueProjectKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("issue with ID %d not found", issueID)
		}
		return fmt.Errorf("failed to fetch issue project key: %v", err)
	}

	if issueProjectKey != sprintProjectKey {
		return fmt.Errorf("issue with ID %d belongs to project %s, which is different from the sprint's project %s", issueID, issueProjectKey, sprintProjectKey)
	}

	// Step 4: Associate the issue with the sprint in the sprint_issues table
	_, err = s.db.Exec("INSERT INTO sprint_issues (sprint_id, issue_id) VALUES (?, ?)", sprintID, issueID)
	if err != nil {
		return fmt.Errorf("failed to add issue to sprint: %v", err)
	}

	return nil
}

func (s *Store) GetIssuesInSprint(sprintID int) ([]types.Issue, error) {
	rows, err := s.db.Query("SELECT id, `key`, summary, `description`, project_key, reporter, assignee, status, issueType FROM issues WHERE sprint_id = ?", sprintID)
	if err != nil {
		return nil, fmt.Errorf("failed to query issues: %v", err)
	}
	defer rows.Close()

	var issues []types.Issue
	for rows.Next() {
		var issue types.Issue
		err := rows.Scan(&issue.ID, &issue.Key, &issue.Summary, &issue.Description, &issue.ProjectKey, &issue.Reporter, &issue.Assignee, &issue.Status, &issue.IssueType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan issue row: %v", err)
		}
		issues = append(issues, issue)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return issues, nil
}
