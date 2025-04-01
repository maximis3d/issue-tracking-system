package types

import (
	"time"
)

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
	ProjectKey  string    `json:"project_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectLead int       `json:"projectLead"`
	IssueCount  int       `json:"issueCount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProjectStore interface {
	GetProjectByKey(name string) (*Project, error)
	GetProjects() ([]Project, error)
	CreateProject(Project) error
}
type ProjectPayload struct {
	ProjectKey  string `json:"project_key" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	ProjectLead int    `json:"projectLead" validate:"required"`
}

type Issue struct {
	ID          int       `json:"id"`
	Summary     string    `json:"summary" validate:"required"`
	Key         string    `json:"key" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ProjectKey  string    `json:"project" validate:"required"`
	Reporter    string    `json:"reporter" validate:"required"`
	Assignee    string    `json:"assignee" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	IssueType   string    `json:"issueType" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

type IssuePayload struct {
	Summary     string `json:"summary" validate:"required"`
	Description string `json:"description" validate:"required"`
	ProjectKey  string `json:"project_key" validate:"required"`
	Reporter    string `json:"reporter" validate:"required"`
	Assignee    string `json:"assignee" validate:"required"`
	Status      string `json:"status" validate:"required"`
	IssueType   string `json:"issueType" validate:"required"`
}
type IssueUpdatePayload struct {
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	ProjectKey  *string `json:"project_key,omitempty"`
	Reporter    *string `json:"reporter,omitempty"`
	Assignee    *string `json:"assignee,omitempty"`
	Status      *string `json:"status,omitempty"`
	IssueType   *string `json:"issueType,omitempty"`
}

type IssueStore interface {
	CreateIssue(issue Issue) error
	UpdateIssue(issue Issue) error
	GetIssueByID(id int) (*Issue, error)
}
