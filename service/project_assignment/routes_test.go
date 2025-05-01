package projectassignment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

func TestProjectAssignmentHandlers(t *testing.T) {
	store := newMockProjectAssignmentStore()
	handler := NewHandler(store)

	t.Run("Assign User", func(t *testing.T) {
		t.Run("should assign user to project", func(t *testing.T) {
			testRequest(t, handler, http.MethodPost, "/projects-assignment/1/assign/2", nil, http.StatusOK)
		})

		t.Run("should fail with invalid projectID", func(t *testing.T) {
			testRequest(t, handler, http.MethodPost, "/projects-assignment/abc/assign/2", nil, http.StatusBadRequest)
		})

		t.Run("should fail with invalid userID", func(t *testing.T) {
			testRequest(t, handler, http.MethodPost, "/projects-assignment/1/assign/xyz", nil, http.StatusBadRequest)
		})
	})

	t.Run("Get Assigned Users", func(t *testing.T) {
		t.Run("should return assigned users", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/projects-assignment/1/assigned-users", nil, http.StatusOK)
		})

		t.Run("should fail with invalid projectID", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/projects-assignment/invalid/assigned-users", nil, http.StatusBadRequest)
		})
	})

	t.Run("Get All Users", func(t *testing.T) {
		testRequest(t, handler, http.MethodGet, "/users", nil, http.StatusOK)
	})

	t.Run("Remove User", func(t *testing.T) {
		t.Run("should remove user if assigned", func(t *testing.T) {
			testRequest(t, handler, http.MethodDelete, "/projects-assignment/1/remove/2", nil, http.StatusOK)
		})

		t.Run("should return 404 if user not assigned", func(t *testing.T) {
			testRequest(t, handler, http.MethodDelete, "/projects-assignment/1/remove/999", nil, http.StatusNotFound)
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

type mockProjectAssignmentStore struct {
	assignments map[int][]types.User
	users       []types.User
}

func newMockProjectAssignmentStore() *mockProjectAssignmentStore {
	return &mockProjectAssignmentStore{
		assignments: map[int][]types.User{
			1: {
				{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com"},
			},
		},
		users: []types.User{
			{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com"},
			{ID: 3, FirstName: "John", LastName: "Smith", Email: "john@example.com"},
		},
	}
}

func (m *mockProjectAssignmentStore) AssignUserToProject(projectID int, userID int, role string) error {
	// Just accept all inputs for simplicity
	m.assignments[projectID] = append(m.assignments[projectID], types.User{ID: userID})
	return nil
}

func (m *mockProjectAssignmentStore) RemoveUserFromProject(projectID int, userID int) error {
	users := m.assignments[projectID]
	newUsers := []types.User{}
	for _, u := range users {
		if u.ID != userID {
			newUsers = append(newUsers, u)
		}
	}
	m.assignments[projectID] = newUsers
	return nil
}

func (m *mockProjectAssignmentStore) GetUsersForProject(projectID int) ([]types.User, error) {
	if users, ok := m.assignments[projectID]; ok {
		return users, nil
	}
	return []types.User{}, nil
}

func (m *mockProjectAssignmentStore) GetAllUsers() ([]types.User, error) {
	return m.users, nil
}

func (m *mockProjectAssignmentStore) IsUserAssignedToProject(projectID int, userID int) (bool, error) {
	users := m.assignments[projectID]
	for _, u := range users {
		if u.ID == userID {
			return true, nil
		}
	}
	return false, nil
}
