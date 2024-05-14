package address

import (
	"database/sql"
	"fmt"
	"github.com/Bilbottom/ecom-application/types"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowsIntoAddress(rows *sql.Rows) (*types.Address, error) {
	address := new(types.Address)
	err := rows.Scan(
		&address.ID,
		&address.UserID,
		&address.IsDefault,
		&address.Line1,
		&address.Line2,
		&address.City,
		&address.Country,
		&address.Postcode,
		&address.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *Store) CreateAddress(address types.Address) (int, error) {
	log.Printf("Creating address for user: %d\n%v", address.UserID, address)
	res, err := s.db.Exec(
		`
		insert into addresses (user_id, is_default, line_1, line_2, city, country, postcode)
		values (?, ?, ?, ?, ?, ?, ?)
		`,
		address.UserID,
		address.IsDefault,
		address.Line1,
		address.Line2,
		address.City,
		address.Country,
		address.Postcode,
	)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (s *Store) GetAddressesByUserID(userId int) ([]types.Address, error) {
	rows, err := s.db.Query(
		`select * from addresses where user_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var address []types.Address
	for rows.Next() {
		a, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, err
		}

		address = append(address, *a)
	}

	return address, nil
}

func (s *Store) GetDefaultAddressByUserID(userId int) (*types.Address, error) {
	rows, err := s.db.Query(
		`select * from addresses where user_id = ? and is_default = 1`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var address []types.Address
	for rows.Next() {
		a, err := scanRowsIntoAddress(rows)
		if err != nil {
			return nil, err
		}

		address = append(address, *a)
	}

	if len(address) != 1 {
		return nil, fmt.Errorf("expected exactly one default address, got %d", len(address))
	}

	return &address[0], nil
}
