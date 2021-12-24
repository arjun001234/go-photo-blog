package controller

import (
	"clean-architecture/entity"
)

type photoController struct {
	userService entity.PhotoService
}

func NewPhotoController(s entity.PhotoService) *photoController {
	return &photoController{s}
}
