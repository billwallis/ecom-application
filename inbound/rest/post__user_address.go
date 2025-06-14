package rest

import (
	"net/http"

	"github.com/billwallis/ecom-application/domain"
)

type CreateAddressPayload struct {
	Line1     string `json:"line1" validate:"required"`
	Line2     string `json:"line2" default:""`
	City      string `json:"city" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Postcode  string `json:"postcode" validate:"required"`
	IsDefault bool   `json:"isDefault" default:"false"`
}

type PostUserAddressHandler struct {
	userAddressUpdater UserAddressUpdater
}

func NewPostUserAddressHandler(userAddressUpdater UserAddressUpdater) *PostUserAddressHandler {
	return &PostUserAddressHandler{
		userAddressUpdater: userAddressUpdater,
	}
}

func (h *PostUserAddressHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var payload CreateAddressPayload

	if err := ParseJSON(request, &payload); err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	userID, err := domain.GetUserIDFromContext(request.Context())
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	addressId, err := h.userAddressUpdater.CreateAddress(domain.Address{
		UserID:    userID,
		IsDefault: payload.IsDefault,
		Line1:     payload.Line1,
		Line2:     payload.Line2,
		City:      payload.City,
		Country:   payload.Country,
		Postcode:  payload.Postcode,
	})
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_ = WriteJSON(writer, http.StatusCreated, map[string]interface{}{
		"address_id": addressId,
	})
}
