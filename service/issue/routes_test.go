package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

const createIssueEndpoint = "/createIssue"

func TestIssueServiceHandlers(t *testing.T) {
	t.Run("Issue Creation", func(t *testing.T) {
		issueStore := newMockIssueStore()
		handler := NewHandler(issueStore)

		t.Run("should fail if payload is invalid", func(t *testing.T) {
			payload := types.Issue{
				ID:          1,
				Summary:     "summary",
				Description: "description",
				ProjectKey:  "project",
				Reporter:    "reporter",
				Assignee:    "assignee",
				Status:      "status",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusBadRequest)
		})

		t.Run("should create a new issue successfully", func(t *testing.T) {
			payload := types.Issue{
				ID:          1,
				Summary:     "summary",
				Description: "description",
				ProjectKey:  "project",
				Reporter:    "reporter",
				Assignee:    "assignee",
				Status:      "status",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusCreated)
		})

		t.Run("should fail if issue already exists", func(t *testing.T) {
			payload := types.Issue{
				ID:          1,
				Summary:     "summary",
				Description: "description",
				ProjectKey:  "project",
				Reporter:    "reporter",
				Assignee:    "assignee",
				Status:      "status",
			}
			testRequest(t, handler, http.MethodPost, createIssueEndpoint, payload, http.StatusBadRequest)
		})
	})
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
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}
}

type mockIssueStore struct {
	issue map[string]types.Issue
}

func newMockIssueStore() *mockIssueStore {
	return &mockIssueStore{
		issue: make(map[string]types.Issue),
	}
}

func (m *mockIssueStore) CreateIssue(issue types.Issue) error {
	idStr := strconv.Itoa(issue.ID) // Convert int ID to string
	if _, exists := m.issue[idStr]; exists {
		return fmt.Errorf("issue with ID %d already exists", issue.ID)
	}
	m.issue[idStr] = issue
	return nil
}
