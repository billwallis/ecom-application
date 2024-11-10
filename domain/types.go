package domain

type Store interface {
	CreateAddress(address Address) (addressId int, err error)
	GetAddressesByUserID(userId int) (addresses []Address, err error)
	GetDefaultAddressByUserID(userId int) (address *Address, err error)

	CreateOrder(order Order) (orderId int, err error)
	CreateOrderItem(orderItem OrderItem) (err error)

	GetProducts() (products []Product, err error)
	GetProductsByIDs(productIds []int) (products []Product, err error)
	UpdateProduct(product Product) (err error)

	GetUserByEmail(email string) (user *User, err error)
	GetUserByID(id int) (user *User, err error)
	CreateUser(user User) (err error)
}
