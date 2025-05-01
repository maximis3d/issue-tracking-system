package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

func TestIssueServiceHandlers(t *testing.T) {
	issueStore := newMockIssueStore()
	handler := NewHandler(issueStore)

	t.Run("Create Issue", func(t *testing.T) {
		t.Run("should fail if payload is invalid", func(t *testing.T) {
			payload := types.IssuePayload{
				Summary:     "",
				Description: "Test Description",
				ProjectKey:  "PRJ",
				Reporter:    "reporter@example.com",
				Assignee:    "assignee@example.com",
				Status:      "open",
				IssueType:   "bug",
			}
			testRequest(t, handler, http.MethodPost, "/createIssue", payload, http.StatusBadRequest)
		})

		t.Run("should create issue successfully", func(t *testing.T) {
			payload := types.IssuePayload{
				Summary:     "Test Issue",
				Description: "Test Description",
				ProjectKey:  "PRJ",
				Reporter:    "reporter@example.com",
				Assignee:    "assignee@example.com",
				Status:      "open",
				IssueType:   "bug",
			}
			testRequest(t, handler, http.MethodPost, "/createIssue", payload, http.StatusCreated)
		})

		t.Run("should fail if issue already exists", func(t *testing.T) {
			payload := types.Issue{
				Summary:     "Test Issue",
				Description: "Test Description",
				ProjectKey:  "PRJ",
				Reporter:    "reporter@example.com",
				Assignee:    "assignee@example.com",
				Status:      "open",
				IssueType:   "bug",
			}

			issueStore.CreateIssue(payload)

			testRequest(t, handler, http.MethodPost, "/createIssue", payload, http.StatusBadRequest)
		})
	})

	t.Run("Get Issue By ID", func(t *testing.T) {

		t.Run("should return the issue successfully", func(t *testing.T) {
			payload := types.Issue{
				ID:          1,
				Summary:     "Test Issue",
				Description: "Test Description",
				ProjectKey:  "PRJ",
				Reporter:    "reporter@example.com",
				Assignee:    "assignee@example.com",
				Status:      "open",
				IssueType:   "bug",
			}

			issueStore.CreateIssue(payload)

			testRequest(t, handler, http.MethodGet, "/issue/1", nil, http.StatusOK)
		})
	})

	t.Run("Get Issues By Project", func(t *testing.T) {
		t.Run("should return 4200 if no issues exist for the project (empty list)", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/issues/PRJ", nil, http.StatusOK)
		})

		t.Run("should return issues for a valid project", func(t *testing.T) {
			payload := types.Issue{
				ID:          1,
				Summary:     "Test Issue",
				Description: "Test Description",
				ProjectKey:  "PRJ",
				Reporter:    "reporter@example.com",
				Assignee:    "assignee@example.com",
				Status:      "open",
				IssueType:   "bug",
			}

			issueStore.CreateIssue(payload)

			testRequest(t, handler, http.MethodGet, "/issues/PRJ", nil, http.StatusOK)
		})
	})

	t.Run("Get Average Cycle Time", func(t *testing.T) {
		t.Run("should return average cycle time for a valid project", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/cycle-time/PRJ", nil, http.StatusOK)
		})
	})

	t.Run("Get Weekly Throughput", func(t *testing.T) {
		t.Run("should return weekly throughput for a valid project", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/throughput/PRJ", nil, http.StatusOK)
		})

	})
}

// testRequest - Helper function to perform HTTP requests and check response
func testRequest(t testing.TB, handler *Handler, method, path string, payload any, expectedStatus int) {
	t.Helper()

	// Marshal the payload into JSON
	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, path, bytes.NewBuffer(marshalled))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Setup the routes for the tests
	router := mux.NewRouter()
	router.HandleFunc("/createIssue", handler.handleCreateIssue).Methods("POST")
	router.HandleFunc("/issue/{id}", handler.handleGetIssueById).Methods("GET")
	router.HandleFunc("/updateIssue/{id}", handler.handleUpdateIssue).Methods("PUT")
	router.HandleFunc("/issues/{key}", handler.handleGetIssuesByProject).Methods("GET")
	router.HandleFunc("/cycle-time/{project_key}", handler.handleGetAverageCycleTime).Methods("GET")
	router.HandleFunc("/throughput/{project_key}", handler.handleGetWeeklyThroughput).Methods("GET")
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}

	if expectedStatus != rr.Code {
		fmt.Printf("Response Body: %s\n", rr.Body.String())
	}
}

// mockIssueStore - Mock implementation of the issue store
type mockIssueStore struct {
	issues map[int]types.Issue
}

func newMockIssueStore() *mockIssueStore {
	return &mockIssueStore{
		issues: make(map[int]types.Issue),
	}
}

func (m *mockIssueStore) CreateIssue(issue types.Issue) error {
	// Simulate conflict if the issue already exists based on Summary and ProjectKey
	for _, existingIssue := range m.issues {
		if existingIssue.Summary == issue.Summary && existingIssue.ProjectKey == issue.ProjectKey {
			return fmt.Errorf("issue already exists")
		}
	}

	if issue.ID == 0 {
		issue.ID = len(m.issues) + 1
	}

	m.issues[issue.ID] = issue
	return nil
}

func (m *mockIssueStore) GetIssueByID(id int) (*types.Issue, error) {
	issue, exists := m.issues[id]
	if !exists {
		return nil, fmt.Errorf("issue not found")
	}
	return &issue, nil
}

func (m *mockIssueStore) GetIssuesByProject(projectKey string) ([]types.Issue, error) {
	var result []types.Issue
	for _, issue := range m.issues {
		if issue.ProjectKey == projectKey {
			result = append(result, issue)
		}
	}
	return result, nil
}

func (m *mockIssueStore) UpdateIssue(issue types.Issue) error {
	if _, exists := m.issues[issue.ID]; !exists {
		return fmt.Errorf("issue not found")
	}
	m.issues[issue.ID] = issue
	return nil
}

func (m *mockIssueStore) GetAverageCycleTime(projectKey string) (time.Duration, error) {
	return time.Duration(0), nil
}

func (m *mockIssueStore) GetWeeklyThroughput(projectKey string) (map[string]int, error) {
	return nil, nil
}
