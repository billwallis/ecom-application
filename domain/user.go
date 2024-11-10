package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserService struct {
	datastore Store
}

func NewUserService(datastore Store) *UserService {
	return &UserService{
		datastore: datastore,
	}
}

func (s *UserService) CreateUser(user User) (err error) {
	return s.datastore.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (user *User, err error) {
	return s.datastore.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (user *User, err error) {
	return s.datastore.GetUserByEmail(email)
}
