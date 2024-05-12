package product

import (
	"database/sql"
	"fmt"
	"github.com/Bilbottom/ecom-application/types"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
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

// GetProducts godoc
//
// @Summary        List accounts
//
// @Description    get all products from the product store
// @Tags           products
// @Accept         json
// @Produce        json
// @Success        200 {array}  []types.Product
// @Router         /products [get]
func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	var products []types.Product
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("select * from products where id in (?%s)", placeholders)

	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var products []types.Product
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec(
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
