package service

import "clean-architecture/entity"

type uservice struct {
	userRepo entity.UserRepository
}

func NewUserService(r entity.UserRepository) *uservice {
	return &uservice{r}
}

func (uservice) CreateUser(u *entity.User) {}

func (uservice) UpdateUser(u *entity.User) {}

func (uservice) DeleteUser(id int64) {}

func (uservice) FindUser(id int64) {}

func (uservice) FindUsers() {}
