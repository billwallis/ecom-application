package rest

import (
	"github.com/Bilbottom/ecom-application/domain"
	"net/http"
)

type HealthChecker interface {
	Check() error
}

type AuthVerifier interface {
	WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc
}

type UserAddressGetter interface {
	GetDefaultAddressByUserID(userId int) (address *domain.Address, err error)
	GetAddressesByUserID(userId int) (addresses []domain.Address, err error)
}

type CartCheckouter interface {
	GetCartItemsIDs(items []domain.CartItem) (productIDs []int, err error)
	CreateOrderFromCart(products []domain.Product, items []domain.CartItem, userID int) (orderID int, totalPrice float64, address string, err error)
}

type UserAddressUpdater interface {
	CreateAddress(address domain.Address) (addressId int, err error)
}

type UserLoginner interface {
	GetUserByEmail(email string) (user *domain.User, err error)
	GetUserByID(id int) (user *domain.User, err error)
}

type UserRegisterer interface {
	CreateUser(user domain.User) (err error)
}

type ProductGetter interface {
	GetProducts() (products []domain.Product, err error)
	GetProductsByIDs(productIds []int) (products []domain.Product, err error)
}

type ProductUpdater interface {
	UpdateProduct(product domain.Product) (err error)
}
