package projectassignment

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

// AssignUserToProject - Assign a user to a project with a role
func (s *Store) AssignUserToProject(projectID int, userID int, role string) error {
	query := `
        INSERT INTO project_assignments (project_id, user_id, role, assigned_at)
        VALUES (?, ?, ?, NOW())
    `
	_, err := s.db.Exec(query, projectID, userID, role)
	return err
}

// RemoveUserFromProject - Remove a user from a project
func (s *Store) RemoveUserFromProject(projectID int, userID int) error {
	query := `
        DELETE FROM project_assignments
        WHERE project_id = ? AND user_id = ?
    `
	_, err := s.db.Exec(query, projectID, userID)
	return err
}

// GetUsersForProject - Get all users assigned to a project
func (s *Store) GetUsersForProject(projectID int) ([]types.User, error) {
	query := `
        SELECT u.id, u.firstName, u.lastName, u.email, u.createdAt
        FROM users u
        JOIN project_assignments pa ON pa.user_id = u.id
        WHERE pa.project_id = ?
    `
	rows, err := s.db.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var u types.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *Store) GetAllUsers() ([]types.User, error) {
	query := `
		SELECT id, firstName, lastName, email
		FROM users
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []types.User
	for rows.Next() {
		var u types.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil

}

// IsUserAssignedToProject - Check if a user is already assigned to a project
func (s *Store) IsUserAssignedToProject(projectID int, userID int) (bool, error) {
	query := `
        SELECT COUNT(*)
        FROM project_assignments
        WHERE project_id = ? AND user_id = ?
    `
	var count int
	err := s.db.QueryRow(query, projectID, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
