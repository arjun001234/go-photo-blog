package repository

import (
	"clean-architecture/entity"
	"database/sql"
)

type photoRepo struct {
	db *sql.DB
}

func NewphotoRepository(d *sql.DB) entity.PhotoRepository {
	return &photoRepo{d}
}

func (r photoRepo) Save(p *entity.Photo) error {
	var err error
	return err
}
func (r photoRepo) Delete(id int64) error {
	var err error
	return err
}
func (r photoRepo) Update(id int64) error {
	var err error
	return err
}
func (r photoRepo) GetOne(id int64) (*entity.Photo, error) {
	p := &entity.Photo{}
	var err error
	return p, err
}
func (r photoRepo) GetAll() (*[]entity.Photo, error) {
	ps := &[]entity.Photo{}
	var err error
	return ps, err
}
