package repository

import (
	"clean-architecture/entity"
	"database/sql"
	"strconv"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(d *sql.DB) entity.UserRepository {
	return &userRepo{d}
}

func (ur *userRepo) Save(u *entity.User) error {

	query := "INSERT INTO PB_USERS(pb_name,pb_email,pb_password) VALUES(\"" + u.Name + "\",\"" + u.Email + "\",\"" + u.Password + "\");"

	smt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}
	r, err := smt.Exec()
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	u.Id = id
	if err != nil {
		return err
	}
	return err
}
func (ur *userRepo) Delete(id int64) error {
	uId := strconv.Itoa(int(id))
	query := `DELETE FROM PB_USERS WHERE id=` + uId + `;`
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}
func (ur *userRepo) Update(u *entity.User) error {
	var err error
	uId := strconv.Itoa(int(u.Id))
	query := `UPDATE PB_USERS SET pb_name="` + u.Name + `", pb_email="` + u.Email + `", pb_password="` + u.Password + `" WHERE id=` + uId + `;`
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}
func (ur *userRepo) GetOne(id int64) (*entity.User, error) {
	u := entity.User{}
	uId := strconv.Itoa(int(id))
	query := `SELECT * FROM PB_USERS WHERE id=` + uId + `;`
	err := ur.db.QueryRow(query).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	return &u, err
}
func (ur *userRepo) GetAll() (*[]entity.User, error) {
	var us []entity.User
	query := `SELECT * FROM PB_USERS;`
	rs, err := ur.db.Query(query)
	for rs.Next() {
		var u entity.User
		err = rs.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return &us, err
		}
		us = append(us, u)
	}
	defer rs.Close()
	return &us, err
}

func (ur *userRepo) GetByEmail(e string) (*entity.User, error) {

	u := entity.User{}

	query := `SELECT * FROM PB_USERS WHERE pb_email="` + e + `";`

	rs, err := ur.db.Query(query)
	if err != nil {
		return &u, err
	}

	for rs.Next() {
		err = rs.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
		break
	}

	return &u, err
}
