package domain_test

import (
	"reflect"
	"testing"

	"github.com/go-sql-driver/mysql"

	"github.com/Bilbottom/ecom-application/db"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/outbound/datastore"
)

const (
	dbUser          = "root"
	dbPassword      = "password"
	dbAddress       = "localhost:3306"
	dbName          = "ecom"
	NetworkProtocol = "tcp"
)

var (
	// TODO: use a test or mock database
	dbConfig = mysql.Config{
		User:                 dbUser,
		Passwd:               dbPassword,
		Addr:                 dbAddress,
		DBName:               dbName,
		Net:                  NetworkProtocol,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	mySQLDataStore, _ = db.NewMySQLStorage(dbConfig)
	dataStore         = datastore.NewStore(mySQLDataStore)

	addressService = domain.NewAddressService(dataStore)
	productService = domain.NewProductService(dataStore)
	orderService   = domain.NewOrderService(dataStore)

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
