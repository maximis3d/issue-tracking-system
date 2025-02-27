package auth

import "testing"

const plainTextPassword string = "testPassword"

func TestHashPassword(t *testing.T) {

	hash, err := HashPassword(plainTextPassword)

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash not to be empty")
	}

	if hash == plainTextPassword {
		t.Errorf("expected hash to be different to plain text password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword(plainTextPassword)

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte(plainTextPassword)) {
		t.Error("expected passwords to match")
	}

	if ComparePassword(hash, []byte("")) {
		t.Error("expected passwords not to match")
	}
}
