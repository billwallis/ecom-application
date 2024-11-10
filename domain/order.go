package domain

import (
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderId"`
	ProductID int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderService struct {
	datastore Store
}

func NewOrderService(datastore Store) *OrderService {
	return &OrderService{
		datastore: datastore,
	}
}

func (s *OrderService) CreateOrder(order Order) (orderId int, err error) {
	return s.datastore.CreateOrder(order)
}

func (s *OrderService) CreateOrderItem(orderItem OrderItem) (err error) {
	return s.datastore.CreateOrderItem(orderItem)
}
