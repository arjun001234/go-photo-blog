package entity

import (
	"time"
)

type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name" validate:"required,alpha,min=4"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Save(u *User) error
	Delete(id int64) error
	Update(u *User) error
	GetOne(id int64) (*User, error)
	GetAll() (*[]User, error)
	GetByEmail(u *User) error
}

type UserService interface {
	CreateUser(u *User) (*Session, error)
	UpdateUser(u *User) error
	DeleteUser(id int64) error
	FindUser(id int64) (*User, error)
	FindUsers() (*[]User, error)
	ValidateUser(u *User) error
	HashPassword(p string) (string, error)
	ComparePassword(p string, hp string) error
	ValidateCredential(u *User) (*Session, error)
	Logout(s *Session) error
	IsUserLoggedIn(s *Session) error
}
