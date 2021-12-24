package controller

import (
	"clean-architecture/entity"
)

type userController struct {
	userService entity.UserService
}

func NewUserController(s entity.UserService) *userController {
	return &userController{s}
}
