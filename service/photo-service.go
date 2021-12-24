package service

import "clean-architecture/entity"

type pservice struct {
	photoRepo entity.PhotoRepository
}

func NewPhotoService(r entity.PhotoRepository) *pservice {
	return &pservice{r}
}

func (pservice) CreatePhoto(u *entity.Photo) {}

func (pservice) UpdatePhoto(u *entity.Photo) {}

func (pservice) DeletePhoto(id int64) {}

func (pservice) FindPhoto(id int64) {}

func (pservice) FindPhotos() {}
