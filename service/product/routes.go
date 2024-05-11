package product

import (
	"github.com/Bilbottom/ecom-application/types"
	"github.com/Bilbottom/ecom-application/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	productStore types.ProductStore
}

func NewHandler(productStore types.ProductStore) *Handler {
	return &Handler{productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost) // TODO: implement payload type, callback
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.productStore.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusOK, ps)
}
