package domain_test

import (
	"testing"

	"github.com/Bilbottom/ecom-application/domain"
)

func Test_AddressCanBeFlattened(t *testing.T) {
	address := domain.Address{
		Line1:    "Bat Cave",
		Line2:    "Wayne Manor",
		City:     "Gotham",
		Country:  "USA",
		Postcode: "12345",
	}
	expected := "Bat Cave\nWayne Manor\nGotham\nUSA\n12345"

	if flat := address.Flatten(); flat != expected {
		t.Errorf("expected %s, got %s", expected, flat)
	}
}
