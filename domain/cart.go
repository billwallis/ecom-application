package domain

import (
	"fmt"
)

type CartItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CartService struct {
	addressService AddressService
	productService ProductService
	orderService   OrderService
}

func NewCartModifier(
	addressService AddressService,
	productService ProductService,
	orderService OrderService,
) *CartService {
	return &CartService{
		addressService: addressService,
		productService: productService,
		orderService:   orderService,
	}
}

func (s *CartService) GetCartItemsIDs(items []CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %d", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (s *CartService) CreateOrderFromCart(ps []Product, items []CartItem, userID int) (int, float64, string, error) {
	// TODO: this logic should be wrapped into a single SQL transaction
	productMap := make(map[int]Product)
	for _, p := range ps {
		productMap[p.ID] = p
	}

	// check if products are in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, "", err
	}

	// calculate total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity in the database
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		if err := s.productService.UpdateProduct(product); err != nil {
			return 0, 0, "", err
		}
	}

	// create order
	defaultAddress, err := s.addressService.GetDefaultAddressByUserID(userID)
	if err != nil {
		return 0, 0, "", err
	}
	address := defaultAddress.Flatten()

	orderID, err := s.orderService.CreateOrder(Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: address,
	})
	if err != nil {
		return 0, 0, "", err
	}

	// create order items
	for _, item := range items {
		err = s.orderService.CreateOrderItem(OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
		if err != nil {
			return 0, 0, "", err
		}
	}

	return orderID, totalPrice, address, nil
}

func checkIfCartIsInStock(cartItems []CartItem, products map[int]Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("the cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []CartItem, products map[int]Product) float64 {
	total := 0.0

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
