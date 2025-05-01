package projectscopes

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

func TestProjectScopeHandlers(t *testing.T) {
	store := newMockScopeStore()
	handler := NewHandler(store)

	t.Run("Create Scope", func(t *testing.T) {
		t.Run("should create scope", func(t *testing.T) {
			scope := types.Scope{Name: "Scope A"}
			testRequest(t, handler, http.MethodPost, "/scopes", scope, http.StatusCreated)
		})

		t.Run("should return 400 on invalid JSON", func(t *testing.T) {
			testRequest(t, handler, http.MethodPost, "/scopes", "{bad json}", http.StatusBadRequest)
		})
	})

	t.Run("Add Project to Scope", func(t *testing.T) {
		t.Run("should add project to scope", func(t *testing.T) {
			body := map[string]string{"project_key": "PROJ1"}
			testRequest(t, handler, http.MethodPost, "/scopes/1", body, http.StatusOK)
		})

		t.Run("should return 400 on invalid scopeID", func(t *testing.T) {
			body := map[string]string{"project_key": "PROJ1"}
			testRequest(t, handler, http.MethodPost, "/scopes/abc", body, http.StatusBadRequest)
		})
	})

	t.Run("Remove Projects from Scope", func(t *testing.T) {
		t.Run("should remove project from scope", func(t *testing.T) {
			body := map[string][]string{"project_keys": {"PROJ1"}}
			testRequest(t, handler, http.MethodDelete, "/scopes/1", body, http.StatusOK)
		})

		t.Run("should return 400 on invalid scopeID", func(t *testing.T) {
			body := map[string][]string{"project_keys": {"PROJ1"}}
			testRequest(t, handler, http.MethodDelete, "/scopes/abc", body, http.StatusBadRequest)
		})
	})

	t.Run("Get Scope Details", func(t *testing.T) {
		t.Run("should get scope details", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/scopes/details/1", nil, http.StatusOK)
		})

		t.Run("should return 400 on invalid scopeID", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/scopes/details/abc", nil, http.StatusBadRequest)
		})
	})

	t.Run("Get Issues by Scope", func(t *testing.T) {
		t.Run("should return issues for scope", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/scopes/issues/1", nil, http.StatusOK)
		})

		t.Run("should return 400 on invalid scopeID", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/scopes/issues/abc", nil, http.StatusBadRequest)
		})
	})

	t.Run("Get All Scopes", func(t *testing.T) {
		t.Run("should return all scopes", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/scopes", nil, http.StatusOK)
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

type mockScopeStore struct {
	scopes map[int]types.Scope
}

func newMockScopeStore() *mockScopeStore {
	return &mockScopeStore{
		scopes: map[int]types.Scope{
			1: {ID: 1, Name: "Scope A"},
		},
	}
}

func (m *mockScopeStore) CreateScope(scope types.Scope) error {
	// Simulate creating a scope
	scope.ID = len(m.scopes) + 1
	m.scopes[scope.ID] = scope
	return nil
}

func (m *mockScopeStore) AddProjectToScope(scopeID int, projectKey string) error {
	// Simulate adding project to scope
	return nil
}

func (m *mockScopeStore) RemoveProjectFromScope(scopeID int, projectKey string) error {
	// Simulate removing project from scope
	return nil
}

func (m *mockScopeStore) GetScopeDetails(scopeID int) (*types.Scope, error) {
	scope, exists := m.scopes[scopeID]
	if !exists {
		return &types.Scope{}, fmt.Errorf("scope not found")
	}
	return &scope, nil
}

func (m *mockScopeStore) GetIssuesByScope(scopeID int) ([]types.Issue, error) {
	// Simulate retrieving issues
	return []types.Issue{}, nil
}

func (m *mockScopeStore) GetAllScopeDetails() ([]types.Scope, error) {
	// Return all mock scopes
	var allScopes []types.Scope
	for _, scope := range m.scopes {
		allScopes = append(allScopes, scope)
	}
	return allScopes, nil
}
