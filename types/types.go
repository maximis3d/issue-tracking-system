package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"="`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectLead string    `json:"projectLead"`
	IssueCount  int       `json:"issueCount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProjectStore interface {
	GetProjectByName(name string) (*Project, error)
	GetProjects() ([]Project, error)
	CreateProject(Project) error
}
type ProjectPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ProjectLead string `json:"projectLead" validate:"required"`
}

type Issue struct {
	Summary     string `json:"summary" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reporter    string `json:"reporter" validate:"required"`
	Assignee    string `json:"assignee" validate:"required"`
	IssueType   string `json:"issueType" validate:"required"`
}
