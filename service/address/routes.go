package address

import (
	"github.com/Bilbottom/ecom-application/service/auth"
	"github.com/Bilbottom/ecom-application/types"
	"github.com/Bilbottom/ecom-application/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	addressStore types.AddressStore
	userStore    types.UserStore
}

func NewHandler(addressStore types.AddressStore, userStore types.UserStore) *Handler {
	return &Handler{
		addressStore: addressStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/address", auth.WithJWTAuth(h.handleCreateAddress, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/address", auth.WithJWTAuth(h.handleGetAddressesByUser, h.userStore)).Methods(http.MethodGet)
}

func (h *Handler) handleCreateAddress(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateAddressPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userID := auth.GetUserIDFromContext(r.Context())
	addressId, err := h.addressStore.CreateAddress(types.Address{
		UserID:    userID,
		IsDefault: payload.IsDefault,
		Line1:     payload.Line1,
		Line2:     payload.Line2,
		City:      payload.City,
		Country:   payload.Country,
		Postcode:  payload.Postcode,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"address_id": addressId,
	})
}

func (h *Handler) handleGetAddressesByUser(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	getDefault := r.URL.Query().Get("default")

	if getDefault == "true" {
		address, err := h.addressStore.GetDefaultAddressByUserID(userID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		_ = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
			"address": address,
		})
	} else {
		addresses, err := h.addressStore.GetAddressesByUserID(userID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		_ = utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
			"addresses": addresses,
		})
	}
}
