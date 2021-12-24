package repository

import (
	"clean-architecture/entity"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) entity.UserRepository {
	return &userRepo{d}
}

func (userRepo) Save(u *entity.User) error {
	var err error
	return err
}
func (userRepo) Delete(id int64) error {
	var err error
	return err
}
func (userRepo) Update(id int64) error {
	var err error
	return err
}
func (userRepo) GetOne(id int64) (*entity.User, error) {
	u := &entity.User{}
	var err error
	return u, err
}
func (userRepo) GetAll() (*[]entity.User, error) {
	// u := &entity.User{}
	us := &[]entity.User{}
	var err error
	return us, err
}
