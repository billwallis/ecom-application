package domain

import (
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"` // TODO: implement a better structure
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductService struct {
	datastore Store
}

func NewProductService(datastore Store) *ProductService {
	return &ProductService{
		datastore: datastore,
	}
}

func (s *ProductService) UpdateProduct(product Product) (err error) {
	return s.datastore.UpdateProduct(product)
}

func (s *ProductService) GetProducts() (products []Product, err error) {
	return s.datastore.GetProducts()
}

func (s *ProductService) GetProductsByIDs(productIds []int) (products []Product, err error) {
	return s.datastore.GetProductsByIDs(productIds)
}
