package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	LastSession time.Time `json:"last_session"`
}

func (u *User) Create() bool {
	return false
}

func (u *User) Read() User {
	return User{Id: u.Id}
}

func (u *User) Update() User {
	return User{Id: u.Id}
}

func (u *User) Delete() bool {
	return false
}
