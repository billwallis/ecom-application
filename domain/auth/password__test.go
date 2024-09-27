package auth_test

import (
	"testing"

	"github.com/Bilbottom/ecom-application/domain/auth"
)

func TestHashPassword(t *testing.T) {
	hash, err := auth.HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be non-empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := auth.HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !auth.ComparePassword(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}

	if auth.ComparePassword(hash, []byte("wrongpassword")) {
		t.Errorf("expected password to not match hash")
	}
}
