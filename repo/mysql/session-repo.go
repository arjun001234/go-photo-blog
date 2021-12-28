package repository

import (
	"clean-architecture/entity"
	"database/sql"
	"strconv"
)

type sessionRepo struct {
	db *sql.DB
}

func NewSessionRepository(d *sql.DB) entity.SessionRepository {
	return &sessionRepo{d}
}

func (sr *sessionRepo) Save(s *entity.Session) error {

	id := strconv.Itoa(int(s.User.Id))
	query := `INSERT INTO PB_SESSIONS(pb_session,pb_user_id) VALUES("` + s.Session + `",` + id + `);`

	smt, err := sr.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = smt.Exec()
	if err != nil {
		return err
	}
	return err
}
func (sr *sessionRepo) Delete(s string) error {
	var err error
	query := `DELETE FROM PB_SESSIONS WHERE pb_session="` + s + `";`
	smt, err := sr.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = smt.Exec()
	return err
}
func (sr *sessionRepo) GetOne(ses string) (*entity.Session, error) {

	s := entity.Session{
		Session: ses,
		User:    &entity.User{},
	}

	query := `SELECT PB_USERS.id,PB_USERS.pb_name,PB_USERS.pb_email,PB_USERS.pb_password,PB_USERS.created_at,PB_USERS.updated_at,PB_SESSIONS.created_at  FROM PB_SESSIONS JOIN PB_USERS ON PB_SESSIONS.pb_user_id = PB_USERS.id WHERE PB_SESSIONS.pb_session = "` + ses + `";`
	err := sr.db.QueryRow(query).Scan(&s.User.Id, &s.User.Name, &s.User.Email, &s.User.Password, &s.User.CreatedAt, &s.User.UpdatedAt, &s.CreatedAt)
	return &s, err
}
