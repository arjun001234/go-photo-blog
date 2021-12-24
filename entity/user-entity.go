package entity

import "time"

type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Save(u *User) error
	Delete(id int64) error
	Update(id int64) error
	GetOne(id int64) (*User, error)
	GetAll() (*[]User, error)
}

type UserService interface {
	CreateUser(u *User)
	UpdateUser(u *User)
	DeleteUser(id int64)
	FindUser(id int64)
	FindUsers()
}
