package rest

import (
	"net/http"

	"github.com/Bilbottom/ecom-application/domain"
)

type GetUserAddressHandler struct {
	userAddressGetter UserAddressGetter
}

func NewGetUserAddressHandler(userAddressGetter UserAddressGetter) *GetUserAddressHandler {
	return &GetUserAddressHandler{
		userAddressGetter: userAddressGetter,
	}
}

func (h *GetUserAddressHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userID, err := domain.GetUserIDFromContext(request.Context())
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	getDefault := request.URL.Query().Get("default")
	if getDefault == "true" {
		address, err := h.userAddressGetter.GetDefaultAddressByUserID(userID)
		if err != nil {
			WriteError(writer, http.StatusInternalServerError, err)
			return
		}

		_ = WriteJSON(writer, http.StatusOK, map[string]interface{}{
			"address": address,
		})
	} else {
		addresses, err := h.userAddressGetter.GetAddressesByUserID(userID)
		if err != nil {
			WriteError(writer, http.StatusInternalServerError, err)
			return
		}

		_ = WriteJSON(writer, http.StatusOK, map[string]interface{}{
			"addresses": addresses,
		})
	}
}
