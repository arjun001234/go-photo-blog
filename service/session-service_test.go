package service

import (
	"clean-architecture/entity"
	repo "clean-architecture/repo/mysql"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"testing"
)

type sessionTestService struct {
	ss entity.SessionService
}

func NewSessionServiceTest(ss entity.SessionService) *sessionTestService {
	return &sessionTestService{ss}
}

func TestSessionService(t *testing.T) {

	s := &entity.Session{
		User: &entity.User{
			Id:       3,
			Name:     "Puig",
			Email:    "puig22@example.com",
			Password: "puig1234",
		},
	}

	db := storage.NewMySQLTestDatabase()
	sr := repo.NewSessionRepository(db)
	ss := NewSessionService(sr)
	ts := NewSessionServiceTest(ss)
	ts.TestSetSession(t, s)
	ts.TestGetSession(t, s)
	ts.TestRemoveSession(t, s)
}

func (ts sessionTestService) TestSetSession(t *testing.T, s *entity.Session) {
	if err := ts.ss.SetSession(s); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Service: Session Added")
}

func (ts sessionTestService) TestGetSession(t *testing.T, s *entity.Session) {
	ns, err := ts.ss.GetUser(s.Session)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if ns.Session != s.Session || ns.User.Id != s.User.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (ts sessionTestService) TestRemoveSession(t *testing.T, s *entity.Session) {
	if err := ts.ss.RemoveSession(s); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Service: Session Removed")
}
