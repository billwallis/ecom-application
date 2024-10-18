package rest

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/domain/auth"
)

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PostUserLoginHandler struct {
	userLoginner UserLoginner
}

func NewPostUserLoginHandler(userLoginner UserLoginner) *PostUserLoginHandler {
	return &PostUserLoginHandler{
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
		errors := err.(validator.ValidationErrors)
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
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

	secret := []byte(config.Envs.JWTSecret)
	token, err := domain.CreateJWT(secret, user.ID)
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_ = WriteJSON(writer, http.StatusOK, map[string]string{"token": token})
}
