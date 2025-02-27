package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT token: %v", err)
	}
	if token == "" {
		t.Error("expected token not to be empty")
	}
}
