package entity

import (
	"time"
)

type User struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name" validate:"required,min=4"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Save(u *User) error                 //tested
	Delete(id int64) error              //tested
	Update(u *User) error               //tested
	GetOne(id int64) (*User, error)     //tested
	GetAll() (*[]User, error)           //tested
	GetByEmail(e string) (*User, error) //tested
}

type UserService interface {
	CreateUser(u *User) (*Session, error)         //tested
	UpdateUser(u *User) error                     //tested
	DeleteUser(id int64) error                    //tested
	FindUser(id int64) (*User, error)             //tested
	FindUsers() (*[]User, error)                  //tested
	ValidateUser(u *User) error                   //tested
	HashPassword(p string) (string, error)        //tested
	ComparePassword(p string, hp string) error    //tested
	ValidateCredential(u *User) (*Session, error) //tested
	Logout(s *Session) error                      //not requied as already tested in session service test
	IsUserLoggedIn(s string) (*Session, error)    //not requied as already tested in session service test
}
