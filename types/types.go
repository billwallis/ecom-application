package types

import (
	"strings"
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ps []int) ([]Product, error)
	UpdateProduct(product Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type AddressStore interface {
	CreateAddress(Address) (int, error)
	GetAddressesByUserID(int) ([]Address, error)
	GetDefaultAddressByUserID(int) (*Address, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"` // TODO: implement a better structure
	CreatedAt   time.Time `json:"createdAt"`
}

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

type Address struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	IsDefault bool      `json:"isDefault"`
	Line1     string    `json:"line1"`
	Line2     string    `json:"line2"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	Postcode  string    `json:"postcode"`
	CreatedAt time.Time `json:"createdAt"`
}

func (a *Address) Flatten() string {
	return strings.Join([]string{
		a.Line1,
		a.Line2,
		a.City,
		a.Country,
		a.Postcode,
	}, "\n")
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=128"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateAddressPayload struct {
	Line1     string `json:"line1" validate:"required"`
	Line2     string `json:"line2" default:""`
	City      string `json:"city" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Postcode  string `json:"postcode" validate:"required"`
	IsDefault bool   `json:"isDefault" default:"false"`
}

type CartItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}
