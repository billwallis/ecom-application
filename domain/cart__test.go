package domain_test

import (
	"reflect"
	"testing"

	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/outbound/datastore"
)

var (
	// TODO: use a test or mock database
	store = newDataStore()

	addressService = domain.NewAddressService(store)
	productService = domain.NewProductService(store)
	orderService   = domain.NewOrderService(store)

	cartService = domain.NewCartService(
		*addressService,
		*productService,
		*orderService,
	)
)

func Test_CartCanGetCartItemsIDs(t *testing.T) {
	cartItems := []domain.CartItem{
		{ProductID: 101, Quantity: 2},
		{ProductID: 102, Quantity: 1},
	}
	expected := []int{101, 102}

	ids, err := cartService.GetCartItemsIDs(cartItems)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(ids, expected) {
		t.Errorf("expected %v, got %v", expected, ids)
	}
}

func newDataStore() domain.Store {
	dbConn, err := config.GetDatabaseConnection(config.NewAppConfig().DBConfig)
	if err != nil {
		panic(err)
	}

	return datastore.NewPostgresStore(dbConn)
}
