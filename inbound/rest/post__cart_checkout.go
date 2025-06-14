package rest

import (
	"fmt"
	"net/http"

	"github.com/billwallis/ecom-application/domain"
)

type CartCheckoutPayload struct {
	Items []domain.CartItem `json:"items" validate:"required"`
}

type PostCartCheckoutHandler struct {
	cartCheckouter CartCheckouter
	productGetter  ProductGetter
}

func NewPostCartCheckoutHandler(cartCheckouter CartCheckouter, productGetter ProductGetter) *PostCartCheckoutHandler {
	return &PostCartCheckoutHandler{
		cartCheckouter: cartCheckouter,
		productGetter:  productGetter,
	}
}

func (h *PostCartCheckoutHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userID, err := domain.GetUserIDFromContext(request.Context())
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	var cart CartCheckoutPayload
	if err := ParseJSON(request, &cart); err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	if err = Validate.Struct(cart); err != nil {
		WriteError(writer, http.StatusBadRequest, fmt.Errorf("invalid payload: %w", err))
		return
	}

	productIDs, err := h.cartCheckouter.GetCartItemsIDs(cart.Items)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	ps, err := h.productGetter.GetProductsByIDs(productIDs)
	if err != nil {
		WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, address, err := h.cartCheckouter.CreateOrderFromCart(ps, cart.Items, userID)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, err)
		return
	}

	_ = WriteJSON(writer, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderID,
		"address":     address,
	})
}
