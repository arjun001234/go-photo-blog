package repository

import (
	"clean-architecture/entity"
	"database/sql"
	"errors"
	"strconv"
)

type photoRepo struct {
	db *sql.DB
}

func NewPhotoRepository(d *sql.DB) entity.PhotoRepository {
	return &photoRepo{d}
}

func (pr *photoRepo) Save(p *entity.Photo) error {

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
func (pr *photoRepo) Delete(id int64) error {
	pId := strconv.Itoa(int(id))
	query := `DELETE FROM PB_PHOTOS WHERE id=` + pId + `;`
	smt, err := pr.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = smt.Exec()
	return err
}

func (pr *photoRepo) GetOne(id int64) (*entity.Photo, error) {
	pId := strconv.Itoa(int(id))
	p := entity.Photo{
		User: &entity.User{},
	}
	query := `SELECT * FROM PB_PHOTOS JOIN PB_USERS ON PB_PHOTOS.pb_user_id = PB_USERS.id WHERE PB_PHOTOS.id=` + pId + `;`
	err := pr.db.QueryRow(query).Scan(&p.Id, &p.Url, &p.User.Id, &p.CreatedAt, &p.User.Id, &p.User.Name, &p.User.Email, &p.User.Password, &p.User.CreatedAt, &p.User.UpdatedAt)
	if err == sql.ErrNoRows {
		return &p, errors.New("photo not found")
	}
	return &p, err
}
func (pr *photoRepo) GetAll() (*[]entity.Photo, error) {
	var ps []entity.Photo
	query := `SELECT * FROM PB_PHOTOS JOIN PB_USERS ON PB_PHOTOS.pb_user_id = PB_USERS.id;`
	r, err := pr.db.Query(query)
	if err != nil {
		return &ps, err
	}
	for r.Next() {
		p := entity.Photo{
			User: &entity.User{},
		}
		err = r.Scan(&p.Id, &p.Url, &p.User.Id, &p.CreatedAt, &p.User.Id, &p.User.Name, &p.User.Email, &p.User.Password, &p.User.CreatedAt, &p.User.UpdatedAt)
		if err != nil {
			return &ps, err
		}
		ps = append(ps, p)
	}
	return &ps, err
}

// func (pr photoRepo) GetByUserAndPhotoId(p *entity.Photo) error {
//    query := `SELECT id,pb_url,created_at FROM PB_PHOTOS WHERE `
// }
