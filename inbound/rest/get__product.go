package rest

import (
	"net/http"
)

type GetProductHandler struct {
	productGetter ProductGetter
}

func NewGetProductHandler(productGetter ProductGetter) *GetProductHandler {
	return &GetProductHandler{
		productGetter: productGetter,
	}
}

func (h *GetProductHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	ps, err := h.productGetter.GetProducts()
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	_ = WriteJSON(writer, http.StatusOK, ps)
}
