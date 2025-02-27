package user

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

func TestUserServiceHandlers(t *testing.T) {
	userStore := newMockUserStore()
	handler := NewHandler(userStore)

	t.Run("User Registration", func(t *testing.T) {
		t.Run("should fail if payload is invalid", func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "user_first_name",
				LastName:  "user_last_name",
				Email:     "invalid_email",
				Password:  "",
			}
			testRequest(t, handler, http.MethodPost, "/register", payload, http.StatusBadRequest)
		})

		t.Run("should register a new user successfully", func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "valid_email@gmail.com",
				Password:  "test123",
			}
			testRequest(t, handler, http.MethodPost, "/register", payload, http.StatusCreated)
		})

		t.Run("should fail if user already exists", func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "valid_email@gmail.com",
				Password:  "test123",
			}
			testRequest(t, handler, http.MethodPost, "/register", payload, http.StatusBadRequest)
		})
	})

	t.Run("User Login", func(t *testing.T) {
		t.Run("should fail with incorrect credentials", func(t *testing.T) {
			payload := types.LoginUserPayload{
				Email:    "invalid_email@gmail.com",
				Password: "wrongpassword",
			}
			testRequest(t, handler, http.MethodPost, "/login", payload, http.StatusBadRequest)
		})

		t.Run("should log in successfully with valid credentials", func(t *testing.T) {
			payload := types.LoginUserPayload{
				Email:    "valid_email@gmail.com",
				Password: "test123",
			}
			testRequest(t, handler, http.MethodPost, "/login", payload, http.StatusOK)
		})
	})
}

// testRequest - Helper function to perform HTTP requests and check response
func testRequest(t testing.TB, handler *Handler, method, path string, payload interface{}, expectedStatus int) {
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
	router.HandleFunc("/register", handler.handleRegister).Methods("POST")
	router.HandleFunc("/login", handler.handleLogin).Methods("POST")
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}
}

// mockUserStore - Mock implementation of the user store
type mockUserStore struct {
	users map[string]types.User
}

func newMockUserStore() *mockUserStore {
	return &mockUserStore{
		users: make(map[string]types.User),
	}
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	user, exists := m.users[email]
	if exists {
		return &user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(user types.User) error {
	if _, exists := m.users[user.Email]; exists {
		return fmt.Errorf("user already exists")
	}
	m.users[user.Email] = user
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
