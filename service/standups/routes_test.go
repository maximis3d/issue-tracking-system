package standups

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

func TestStandupHandlers(t *testing.T) {
	store := newMockStandupStore()
	handler := NewHandler(store)

	t.Run("Create Standup", func(t *testing.T) {
		t.Run("should create a new standup", func(t *testing.T) {
			standup := types.Standup{
				ProjectKey: "project-1",
			}
			testRequest(t, handler, http.MethodPost, "/standups/start", standup, http.StatusCreated)
		})
	})

	t.Run("End Standup", func(t *testing.T) {
		t.Run("should end an active standup", func(t *testing.T) {
			standup := types.Standup{
				ProjectKey: "project-1",
			}
			testRequest(t, handler, http.MethodPost, "/standups/end", standup, http.StatusOK)
		})
	})
}

func testRequest(t testing.TB, handler *Handler, method, path string, payload any, expectedStatus int) {
	t.Helper()

	var body *bytes.Buffer
	if payload != nil {
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		body = bytes.NewBuffer(marshalled)
	} else {
		body = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	handler.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}
}

// -------------------- MOCK STORE --------------------

type mockStandupStore struct{}

func newMockStandupStore() *mockStandupStore {
	return &mockStandupStore{}
}

func (m *mockStandupStore) CreateStandup(standup types.Standup) error {
	// Simulating standup creation
	return nil
}

func (m *mockStandupStore) EndCurrentStandUp(standup types.Standup) error {
	// Simulating ending a standup
	return nil
}

func (m *mockStandupStore) GetActiveStandup(standup types.Standup) (*types.Standup, error) {
	// Simulate an active standup
	return &types.Standup{ProjectKey: "project-1"}, nil
}

func (m *mockStandupStore) FilterTickets(project types.Project) ([]types.Issue, error) {
	// Simulating issue filtering
	return []types.Issue{
		{ID: 1, Summary: "Issue 1", ProjectKey: project.ProjectKey},
		{ID: 2, Summary: "Issue 2", ProjectKey: project.ProjectKey},
	}, nil
}

func (m *mockStandupStore) FilterTicketsByEndTime(projectKey string, lastEndTime sql.NullTime) ([]types.Issue, error) {
	// Simulating issue filtering by end time
	return []types.Issue{
		{ID: 1, Summary: "Issue 1", ProjectKey: projectKey},
		{ID: 2, Summary: "Issue 2", ProjectKey: projectKey},
	}, nil
}

func (m *mockStandupStore) GetLastStandupEndTime(projectKey string) (sql.NullTime, error) {
	// Simulating fetching the last standup end time
	return sql.NullTime{Valid: true, Time: time.Now()}, nil
}
