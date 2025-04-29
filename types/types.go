package types

import (
	"database/sql"
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
	WIPLimit    int       `json:"wip_limit"`
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
	WIPLimit    int    `json:"wip_limit" validate:"required"`
}

type Issue struct {
	ID          int    `json:"id"`
	Summary     string `json:"summary" validate:"required"`
	Key         string `json:"key" validate:"required"`
	Description string `json:"description" validate:"required"`
	ProjectKey  string `json:"project" validate:"required"`
	Reporter    string `json:"reporter" validate:"required"`
	Assignee    string `json:"assignee" validate:"required"`
	Status      string `json:"status" validate:"required"`
	IssueType   string `json:"issueType" validate:"required"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
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
	GetIssuesByProject(projectKey string) ([]Issue, error)
}

type Standup struct {
	ID         int        `json:"id"`
	ProjectKey string     `json:"project_key" validate:"required"`
	StartTime  time.Time  `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	CreatedAt  time.Time  `json:"created_at"`
}

type StandupStore interface {
	CreateStandup(Standup) error
	EndCurrentStandUp(Standup) error
	GetActiveStandup(Standup) (*Standup, error)
	FilterTickets(Project) ([]Issue, error)
	FilterTicketsByEndTime(projectKey string, lastEndTime sql.NullTime) ([]Issue, error)
	GetLastStandupEndTime(projectKey string) (sql.NullTime, error)
}

type Scope struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Projects    []string `json:"projects"`
}

type ScopeStore interface {
	CreateScope(Scope) error
	AddProjectToScope(scopeID int, projectKey string) error
	GetIssuesByScope(scopeID int) ([]Issue, error)
}

type ProjectAssignment struct {
	ProjectID  int       `json:"project_id"`
	UserID     int       `json:"user_id"`
	Role       string    `json:"role"`
	AssignedAt time.Time `json:"assigned_at"`
}

type ProjectAssignmentStore interface {
	AssignUserToProject(projectID int, userID int, role string) error
	RemoveUserFromProject(projectID int, userID int) error
	GetUsersForProject(projectID int) ([]User, error)
	GetAllUsers() ([]User, error)
}
