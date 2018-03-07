package models

import (
	"time"

	"github.com/ken-aio/go-echo-sqlboiler/app/infra/db"
	"github.com/volatiletech/sqlboiler/queries/qm"
	null "gopkg.in/volatiletech/null.v6"
)

// UserCreate new user
type UserCreate struct {
	ID uint64 `json:"id"`
}

// UserList list userr struct
type UserList struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// User basic user struct
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}

// Create insert into user
func (u *User) Create() (uint64, error) {
	user := &db.User{
		Name:      u.Name,
		Birthdate: null.Time{u.Birthdate, true},
		Gender:    u.Gender,
	}
	err := user.InsertG()
	return user.ID, err
}

// List select user list
func (u *User) List() ([]*UserList, error) {
	users, err := db.UsersG(qm.Select("id", "name"), qm.OrderBy("ID asc")).All()
	if err != nil {
		return nil, err
	}

	list := make([]*UserList, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		list[i] = &UserList{user.ID, user.Name}
	}
	return list, nil
}
