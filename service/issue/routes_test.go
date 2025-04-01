package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

const createIssueEndpoint = "/createIssue"
const updateIssueEndpoint = "/updateIssue/{id}"

func TestIssueServiceHandlers(t *testing.T) {
	t.Run("Issue Creation", func(t *testing.T) {
		issueStore := newMockIssueStore()
		handler := NewHandler(issueStore)

		t.Run("should fail if payload is invalid", func(t *testing.T) {
			payload := map[string]string{
				"invalid": "data",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusBadRequest)
		})

		t.Run("should create a new issue successfully", func(t *testing.T) {
			payload := types.IssuePayload{
				Summary:     "summary",
				Description: "description",
				ProjectKey:  "project",
				Reporter:    "reporter",
				Assignee:    "assignee",
				Status:      "status",
				IssueType:   "bug",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusCreated)
		})

		t.Run("should fail if issue already exists", func(t *testing.T) {
			payload := types.IssuePayload{
				Summary:     "summary",
				Description: "description",
				ProjectKey:  "project",
				Reporter:    "reporter",
				Assignee:    "assignee",
				Status:      "status",
				IssueType:   "bug",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusBadRequest)
		})
	})

	t.Run("Issue Update", func(t *testing.T) {
		issueStore := newMockIssueStore()
		handler := NewHandler(issueStore)

		// Create an issue first
		issueStore.CreateIssue(types.Issue{
			ID:          1,
			Summary:     "Original summary",
			Description: "Original description",
			ProjectKey:  "project",
			Reporter:    "reporter",
			Assignee:    "assignee",
			Status:      "open",
			IssueType:   "bug",
		})

		t.Run("should update an existing issue", func(t *testing.T) {
			payload := types.IssueUpdatePayload{
				Summary:     strPtr("Updated summary"),
				Description: strPtr("Updated description"),
			}
			testRequest(t, handler, http.MethodPut, "/updateIssue/1", payload, http.StatusOK)
		})

		t.Run("should return error if issue not found", func(t *testing.T) {
			payload := types.IssueUpdatePayload{
				Summary: strPtr("New summary"),
			}
			testRequest(t, handler, http.MethodPut, "/updateIssue/999", payload, http.StatusNotFound)
		})
	})
}

// Helper function to create a pointer to a string
func strPtr(s string) *string {
	return &s
}

func testRequest(t testing.TB, handler *Handler, method, path string, payload any, expectedStatus int) {
	t.Helper()

	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(marshalled))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(createIssueEndpoint, handler.handleCreateIssue).Methods("POST")
	router.HandleFunc(updateIssueEndpoint, handler.handleUpdateIssue).Methods("PUT")
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}
}

type mockIssueStore struct {
	issues map[int]types.Issue
}

func newMockIssueStore() *mockIssueStore {
	return &mockIssueStore{
		issues: make(map[int]types.Issue),
	}
}

func (m *mockIssueStore) CreateIssue(issue types.Issue) error {
	if _, exists := m.issues[issue.ID]; exists {
		return fmt.Errorf("issue with ID %d already exists", issue.ID)
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

func (m *mockIssueStore) UpdateIssue(issue types.Issue) error {
	if _, exists := m.issues[issue.ID]; !exists {
		return fmt.Errorf("issue not found")
	}
	m.issues[issue.ID] = issue
	return nil
}
