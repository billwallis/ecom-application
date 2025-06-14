package rest

import (
	"fmt"
	"net/http"

	"github.com/billwallis/ecom-application/config"
	"github.com/billwallis/ecom-application/domain"
	"github.com/billwallis/ecom-application/domain/auth"
)

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PostUserLoginHandler struct {
	appConfig    config.AppConfig
	authService  domain.AuthService
	userLoginner UserLoginner
}

func NewPostUserLoginHandler(
	appConfig config.AppConfig,
	authService domain.AuthService,
	userLoginner UserLoginner,
) *PostUserLoginHandler {
	return &PostUserLoginHandler{
		appConfig:    appConfig,
		authService:  authService,
		userLoginner: userLoginner,
	}
}

func (h *PostUserLoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var payload LoginUserPayload
	if err := ParseJSON(request, &payload); err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload %w", err))
		return
	}

	user, err := h.userLoginner.GetUserByEmail(payload.Email)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(user.Password, []byte(payload.Password)) {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte(h.appConfig.AuthConfig.JWTSecret)
	token, err := h.authService.CreateJWT(secret, user.ID)
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_ = WriteJSON(writer, http.StatusOK, map[string]string{"token": token})
}
