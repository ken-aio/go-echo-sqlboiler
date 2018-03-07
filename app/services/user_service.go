package services

import (
	"time"

	"github.com/ken-aio/go-echo-sqlboiler/app/models"
)

// UserCreate new user create
func UserCreate(name string, birthdate time.Time, gender string) (uint64, error) {
	u := &models.User{Name: name, Birthdate: birthdate, Gender: gender}
	return u.Create()
}

// UserList list user
func UserList() ([]*models.UserList, error) {
	u := &models.User{}
	return u.List()
}
