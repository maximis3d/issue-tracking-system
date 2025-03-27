package project

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

func TestProjectHandlers(t *testing.T) {
	projectStore := newMockProjectStore()
	handler := NewHandler(projectStore)

	t.Run("Get Projects", func(t *testing.T) {
		t.Run("should return empty list if no projects exist", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/projects", nil, http.StatusNotFound)
		})

		t.Run("should return list of projects", func(t *testing.T) {
			projectStore.CreateProject(types.Project{Name: "Test Project", Description: "Test Desc", ProjectLead: "Alice"})
			testRequest(t, handler, http.MethodGet, "/projects", nil, http.StatusOK)
		})
	})

	t.Run("Get Project By Name", func(t *testing.T) {
		t.Run("should return 404 if project does not exist", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/projects/Unknown", nil, http.StatusNotFound)
		})

		t.Run("should return project if it exists", func(t *testing.T) {
			testRequest(t, handler, http.MethodGet, "/projects/Test Project", nil, http.StatusOK)
		})
	})

	t.Run("Create Project", func(t *testing.T) {
		t.Run("should fail if payload is invalid", func(t *testing.T) {
			payload := map[string]string{"Name": ""} // Missing required fields
			testRequest(t, handler, http.MethodPost, "/projects", payload, http.StatusBadRequest)
		})

		t.Run("should create project successfully", func(t *testing.T) {
			payload := types.ProjectPayload{Name: "New Project", Description: "Desc", ProjectLead: "Lead"}
			testRequest(t, handler, http.MethodPost, "/projects", payload, http.StatusCreated)
		})

		t.Run("should fail if project already exists", func(t *testing.T) {
			payload := types.ProjectPayload{Name: "New Project", Description: "Desc", ProjectLead: "Lead"}
			testRequest(t, handler, http.MethodPost, "/projects", payload, http.StatusBadRequest)
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

// mockProjectStore - Mock implementation of the project store
type mockProjectStore struct {
	projects map[string]types.Project
}

func newMockProjectStore() *mockProjectStore {
	return &mockProjectStore{
		projects: make(map[string]types.Project),
	}
}

func (m *mockProjectStore) GetProjects() ([]types.Project, error) {
	if len(m.projects) == 0 {
		return nil, fmt.Errorf("no projects found")
	}
	var projects []types.Project
	for _, project := range m.projects {
		projects = append(projects, project)
	}
	return projects, nil
}

func (m *mockProjectStore) GetProjectByName(name string) (*types.Project, error) {
	project, exists := m.projects[name]
	if !exists {
		return nil, fmt.Errorf("project not found")
	}
	return &project, nil
}

func (m *mockProjectStore) CreateProject(project types.Project) error {
	if _, exists := m.projects[project.Name]; exists {
		return fmt.Errorf("project already exists")
	}
	m.projects[project.Name] = project
	return nil
}
