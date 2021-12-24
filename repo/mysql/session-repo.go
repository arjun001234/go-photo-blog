package repository

import (
	"clean-architecture/entity"
	"database/sql"
)

type sessionRepo struct {
	db *sql.DB
}

func NewsessionRepository(d *sql.DB) entity.SessionRepository {
	return &sessionRepo{d}
}

func (sessionRepo) Save(s *entity.Session) error {
	var err error
	return err
}
func (sessionRepo) Delete(id int64) error {
	var err error
	return err
}
func (sessionRepo) GetOne(id int64) (*entity.Session, error) {
	s := &entity.Session{}
	var err error
	return s, err
}
