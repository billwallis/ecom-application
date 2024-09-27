package rest

import (
	"net/http"
)

type PostProductHandler struct {
	productUpdater ProductUpdater
}

func NewPostProductHandler(productUpdater ProductUpdater) *PostProductHandler {
	return &PostProductHandler{
		productUpdater: productUpdater,
	}
}

func (h *PostProductHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_ = WriteJSON(writer, http.StatusOK, nil)
}
