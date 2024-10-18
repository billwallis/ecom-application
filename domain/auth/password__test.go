package auth_test

import (
	"testing"

	"github.com/Bilbottom/ecom-application/domain/auth"
)

const (
	plainPassword = "password"
)

func Test_PasswordCanBeHashed(t *testing.T) {
	hash, err := auth.HashPassword(plainPassword)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be non-empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to be different from plain")
	}
}

func Test_PasswordsCanBeCompared(t *testing.T) {
	hash, _ := auth.HashPassword(plainPassword)

	if !auth.ComparePassword(hash, []byte(plainPassword)) {
		t.Errorf("expected password to match hash")
	}

	if auth.ComparePassword(hash, []byte("wrongpassword")) {
		t.Errorf("expected password to not match hash")
	}
}
