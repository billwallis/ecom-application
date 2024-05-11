package order

import (
	"database/sql"
	"github.com/Bilbottom/ecom-application/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec(
		`
		insert into orders(user_id, total, status, address)
		values (?, ?, ?, ?)
		`,
		order.UserID, order.Total, order.Status, order.Address,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec(
		`
		insert into order_items(order_id, product_id, quantity, price)
		values (?, ?, ?, ?)
		`,
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price,
	)
	return err
}
