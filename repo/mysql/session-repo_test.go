package repository

import (
	"clean-architecture/entity"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

type sessionRepoTest struct {
	sr entity.SessionRepository
}

func NewSessionTestRepo(sr entity.SessionRepository) *sessionRepoTest {
	return &sessionRepoTest{sr}
}

func TestSessionRepository(t *testing.T) {

	s := &entity.Session{
		Session: uuid.New().String(),
		User: &entity.User{
			Id:       1,
			Name:     "messi",
			Email:    "messi34@example.com",
			Password: "messi1234",
		},
	}

	db := storage.NewMySQLTestDatabase()
	sr := NewSessionRepository(db)
	tr := NewSessionTestRepo(sr)
	tr.TestSave(t, s)
	tr.TestDelete(t, s)
}

func (tr *sessionRepoTest) TestSave(t *testing.T, s *entity.Session) {
	if err := tr.sr.Save(s); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("New Session Saved")
}

func (tr *sessionRepoTest) TestGetOne(t *testing.T, s *entity.Session) {
	ss, err := tr.sr.GetOne(s.Session)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if ss.Session != s.Session || ss.User.Id != s.User.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (tr *sessionRepoTest) TestDelete(t *testing.T, s *entity.Session) {
	if err := tr.sr.Delete(s.Session); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Session Deleted")
}
