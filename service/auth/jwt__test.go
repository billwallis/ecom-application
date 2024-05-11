package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")
	userID := 1

	token, err := CreateJWT(secret, userID)
	if err != nil {
		t.Errorf("error creating jwt: %v", err)
	}

	if token == "" {
		t.Errorf("expected token to be non-empty")
	}
}
