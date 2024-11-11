package rest

import (
	"fmt"
	"net/http"

	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/domain/auth"
)

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=128"`
}

type PostUserRegisterHandler struct {
	userRegisterer UserRegisterer
	userLoginner   UserLoginner
}

func NewPostUserRegisterHandler(
	userRegisterer UserRegisterer,
	userLoginner UserLoginner,
) *PostUserRegisterHandler {
	return &PostUserRegisterHandler{
		userRegisterer: userRegisterer,
		userLoginner:   userLoginner,
	}
}

func (h *PostUserRegisterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var payload RegisterUserPayload
	if err := ParseJSON(request, &payload); err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload %w", err))
		return
	}

	_, err := h.userLoginner.GetUserByEmail(payload.Email)
	if err == nil {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	err = h.userRegisterer.CreateUser(domain.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_ = WriteJSON(writer, http.StatusCreated, nil)
}
