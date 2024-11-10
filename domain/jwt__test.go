package domain_test

import (
	"testing"

	"github.com/Bilbottom/ecom-application/domain"
)

func Test_JWTCanBeCreated(t *testing.T) {
	secret := []byte("secret")
	userID := 1

	token, err := domain.CreateJWT(secret, userID)
	if err != nil {
		t.Errorf("error creating jwt: %v", err)
	}
	if token == "" {
		t.Errorf("expected token to be non-empty")
	}
}
