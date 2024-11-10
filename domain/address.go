package domain

import (
	"strings"
	"time"
)

type Address struct {
	ID        int
	UserID    int
	IsDefault bool
	Line1     string
	Line2     string
	City      string
	Country   string
	Postcode  string
	CreatedAt time.Time
}

func (a *Address) Flatten() string {
	return strings.Join(
		[]string{
			a.Line1,
			a.Line2,
			a.City,
			a.Country,
			a.Postcode,
		},
		"\n",
	)
}

type AddressService struct {
	datastore Store
}

func NewAddressService(datastore Store) *AddressService {
	return &AddressService{
		datastore: datastore,
	}
}

func (s *AddressService) CreateAddress(address Address) (addressId int, err error) {
	return s.datastore.CreateAddress(address)
}

func (s *AddressService) GetAddressesByUserID(userId int) (addresses []Address, err error) {
	return s.datastore.GetAddressesByUserID(userId)
}

func (s *AddressService) GetDefaultAddressByUserID(userId int) (address *Address, err error) {
	return s.datastore.GetDefaultAddressByUserID(userId)
}
