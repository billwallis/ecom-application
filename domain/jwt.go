package domain

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/billwallis/ecom-application/config"
)

type contextKey string

const UserKey contextKey = "userID"

type AuthService struct {
	authConfig  config.AuthConfig
	userService UserService
}

func NewAuthService(authConfig config.AuthConfig, userService UserService) *AuthService {
	return &AuthService{
		authConfig:  authConfig,
		userService: userService,
	}
}

func (as *AuthService) WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)

		token, err := as.validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			//permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			//permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)

		u, err := as.userService.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			//permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

//func permissionDenied(w http.ResponseWriter) {
//	rest.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
//}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func (as *AuthService) validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(as.authConfig.JWTSecret), nil
	})
}

func (as *AuthService) CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(as.authConfig.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1, errors.Errorf("user id not found")
	}
	return userID, nil
}
