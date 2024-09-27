package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Bilbottom/ecom-application/domain"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateAddress(address domain.Address) (addressId int, err error) {
	log.Printf("Creating address for user: %d\n%v", address.UserID, address)
	res, err := s.db.Exec(
		`
		insert into addresses (user_id, is_default, line_1, line_2, city, country, postcode)
		values (?, ?, ?, ?, ?, ?, ?)
		`,
		address.UserID,
		address.IsDefault,
		address.Line1,
		address.Line2,
		address.City,
		address.Country,
		address.Postcode,
	)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (s *Store) GetAddressesByUserID(userId int) (addresses []domain.Address, err error) {
	rows, err := s.db.Query(
		`select * from addresses where user_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var address []domain.Address
	for rows.Next() {
		a, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, err
		}

		address = append(address, *a)
	}

	return address, nil
}

func (s *Store) GetDefaultAddressByUserID(userId int) (address *domain.Address, err error) {
	rows, err := s.db.Query(
		`select * from addresses where user_id = ? and is_default = 1`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var addresses []domain.Address
	for rows.Next() {
		addr, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *addr)
	}

	if len(addresses) != 1 {
		return nil, fmt.Errorf("expected exactly one default address, got %d", len(addresses))
	}

	return &addresses[0], nil
}

func scanRowsIntoAddress(rows *sql.Rows) (*domain.Address, error) {
	address := new(domain.Address)
	err := rows.Scan(
		&address.ID,
		&address.UserID,
		&address.IsDefault,
		&address.Line1,
		&address.Line2,
		&address.City,
		&address.Country,
		&address.Postcode,
		&address.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Store) CreateOrder(order domain.Order) (orderId int, err error) {
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

func (s *Store) CreateOrderItem(orderItem domain.OrderItem) (err error) {
	_, err = s.db.Exec(
		`
		insert into order_items(order_id, product_id, quantity, price)
		values (?, ?, ?, ?)
		`,
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price,
	)
	return err
}

func (s *Store) GetProducts() (products []domain.Product, err error) {
	rows, err := s.db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByIDs(productIds []int) (products []domain.Product, err error) {
	placeholders := strings.Repeat(",?", len(productIds)-1)
	query := fmt.Sprintf("select * from products where id in (?%s)", placeholders)

	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*domain.Product, error) {
	product := new(domain.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) UpdateProduct(product domain.Product) (err error) {
	_, err = s.db.Exec(
		`
		update products
		set name = ?,
			price = ?,
			image = ?,
			description = ?,
			quantity = ?
		where id = ?
		`,
		product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (user *domain.User, err error) {
	rows, err := s.db.Query("select * from users where email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(domain.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (user *domain.User, err error) {
	rows, err := s.db.Query("select * from users where id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(domain.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*domain.User, error) {
	user := new(domain.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateUser(user domain.User) (err error) {
	_, err = s.db.Exec(
		"insert into users (first_name, last_name, email, password) values (?, ?, ?, ?)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
