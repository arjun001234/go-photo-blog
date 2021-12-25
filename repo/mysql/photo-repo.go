package repository

import (
	"clean-architecture/entity"
	"database/sql"
	"strconv"
)

type photoRepo struct {
	db *sql.DB
}

func NewphotoRepository(d *sql.DB) entity.PhotoRepository {
	return &photoRepo{d}
}

func (pr photoRepo) Save(p *entity.Photo) error {

	id := strconv.Itoa(int(p.User.Id))

	query := `INSERT INTO PB_PHOTOS(pb_url,pb_user_id) VALUES("` + p.Url + `",` + id + `);`

	smt, err := pr.db.Prepare(query)
	if err != nil {
		return err
	}

	result, err := smt.Exec()
	if err != nil {
		return err
	}

	pId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = pId

	return err
}
func (pr photoRepo) Delete(id int64) error {
	pId := strconv.Itoa(int(id))
	query := `DELETE FROM PB_PHOTOS WHERE id=` + pId + `;`
	smt, err := pr.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = smt.Exec()
	return err
}

func (pr photoRepo) GetOne(id int64) (*entity.Photo, error) {
	var p entity.Photo
	// pId := strconv.Itoa(int(id))

	// pr.db.QueryRow()
	var err error
	return &p, err
}
func (r photoRepo) GetAll() (*[]entity.Photo, error) {
	ps := &[]entity.Photo{}
	var err error
	return ps, err
}
