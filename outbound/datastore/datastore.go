package datastore

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"

	"github.com/Bilbottom/ecom-application/domain"
)

type PostgresStore struct {
	dbConn *pgx.Conn
}

func NewPostgresStore(dbConn *pgx.Conn) *PostgresStore {
	return &PostgresStore{
		dbConn: dbConn,
	}
}

func (ps *PostgresStore) CreateAddress(address domain.Address) (addressId int, err error) {
	log.Printf("Creating address for user: %d\n%v", address.UserID, address)
	res := ps.dbConn.QueryRow(
		context.Background(),
		`
		insert into ecom.addresses (user_id, is_default, line_1, line_2, city, country, postcode)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning id
		`,
		address.UserID,
		address.IsDefault,
		address.Line1,
		address.Line2,
		address.City,
		address.Country,
		address.Postcode,
	)

	var lastInsertId int
	err = res.Scan(&lastInsertId)
	if err != nil {
		return -1, fmt.Errorf("failed to create address: %w", err)
	}

	return lastInsertId, nil
}

func (ps *PostgresStore) GetAddressesByUserID(userId int) (addresses []domain.Address, err error) {
	rows, err := ps.dbConn.Query(
		context.Background(),
		`select * from ecom.addresses where user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get addresses: %w", err)
	}

	var address []domain.Address
	for rows.Next() {
		a, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows into address: %w", err)
		}

		address = append(address, *a)
	}

	return address, nil
}

func (ps *PostgresStore) GetDefaultAddressByUserID(userId int) (address *domain.Address, err error) {
	rows, err := ps.dbConn.Query(
		context.Background(),
		`select * from ecom.addresses where user_id = $1 and is_default`,
		userId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get default address: %w", err)
	}

	var addresses []domain.Address
	for rows.Next() {
		addr, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows into address: %w", err)
		}

		addresses = append(addresses, *addr)
	}

	if len(addresses) != 1 {
		return nil, fmt.Errorf("expected exactly one default address, got %d", len(addresses))
	}

	return &addresses[0], nil
}

func scanRowsIntoAddress(rows pgx.Rows) (*domain.Address, error) {
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

func (ps *PostgresStore) CreateOrder(order domain.Order) (orderId int, err error) {
	res := ps.dbConn.QueryRow(
		context.Background(),
		`
		insert into ecom.orders(user_id, total, status, address)
		values ($1, $2, $3, $4)
		returning id
		`,
		order.UserID, order.Total, order.Status, order.Address,
	)

	var lastInsertId int
	err = res.Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}
	log.Printf("Created order with ID %d", lastInsertId)

	return lastInsertId, nil
}

func (ps *PostgresStore) CreateOrderItem(orderItem domain.OrderItem) (err error) {
	_, err = ps.dbConn.Exec(
		context.Background(),
		`
		insert into ecom.order_items(order_id, product_id, quantity, price)
		values ($1, $2, $3, $4)
		`,
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price,
	)

	if err != nil {
		return fmt.Errorf("failed to create order item: %w", err)
	}
	return nil
}

func (ps *PostgresStore) GetProducts() (products []domain.Product, err error) {
	rows, err := ps.dbConn.Query(
		context.Background(),
		"select * from ecom.products",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows into product: %w", err)
		}

		products = append(products, *p)
	}

	return products, nil
}

func (ps *PostgresStore) GetProductsByIDs(productIds []int) (products []domain.Product, err error) {
	if len(productIds) == 0 {
		return []domain.Product{}, nil
	}

	placeholders := "$1"
	for i := 2; i <= len(productIds); i++ {
		placeholders += fmt.Sprintf(",$%d", i)
	}
	query := fmt.Sprintf("select * from ecom.products where id in (%s)", placeholders)

	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := ps.dbConn.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by ids: %w", err)
	}

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows into product: %w", err)
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows pgx.Rows) (*domain.Product, error) {
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

func (ps *PostgresStore) UpdateProduct(product domain.Product) (err error) {
	_, err = ps.dbConn.Exec(
		context.Background(),
		`
		update ecom.products
		set name = $1,
			price = $2,
			image = $3,
			description = $4,
			quantity = $5
		where id = $6
		`,
		product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (ps *PostgresStore) GetUserByEmail(email string) (user *domain.User, err error) {
	rows, err := ps.dbConn.Query(
		context.Background(),
		"select * from ecom.users where email = $1",
		email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	u := new(domain.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into user: %w", err)
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (ps *PostgresStore) GetUserByID(id int) (user *domain.User, err error) {
	rows, err := ps.dbConn.Query(
		context.Background(),
		"select * from ecom.users where id = $1",
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	u := new(domain.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row into user: %w", err)
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows pgx.Rows) (*domain.User, error) {
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

func (ps *PostgresStore) CreateUser(user domain.User) (err error) {
	_, err = ps.dbConn.Exec(
		context.Background(),
		"insert into ecom.users (first_name, last_name, email, password) values ($1, $2, $3, $4)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
