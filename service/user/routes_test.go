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
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// route constants
	const registerConst = "/register"

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user_first_name",
			LastName:  "user_last_name",
			Email:     "invalid_email",
			Password:  "",
		}

		rr := performrequest(t, handler, http.MethodGet, registerConst, payload)

		assertCorrectMessage(t, rr, http.StatusBadRequest)
	})

	t.Run("should correctly register the new user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user_first_name",
			LastName:  "user_last_name",
			Email:     "valid_email@gmail.com",
			Password:  "test123",
		}

		rr := performrequest(t, handler, http.MethodGet, registerConst, payload)

		assertCorrectMessage(t, rr, http.StatusCreated)
	})
}

func performrequest(t testing.TB, handler *Handler, method string, path string, payload interface{}) *httptest.ResponseRecorder {
	t.Helper()

	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)

	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(marshalled))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc(path, handler.handleRegister)
	router.ServeHTTP(rr, req)

	return rr
}

func assertCorrectMessage(t testing.TB, rr *httptest.ResponseRecorder, expectedResponseCode int) {
	t.Helper()

	if rr.Code != expectedResponseCode {
		t.Errorf("expected status code %d, got %d", expectedResponseCode, rr.Code)
	}
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
