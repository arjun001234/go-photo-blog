package service

import (
	"clean-architecture/entity"

	"github.com/google/uuid"
)

type sservice struct {
	sessionRepo entity.SessionRepository
}

func NewSessionService(r entity.SessionRepository) entity.SessionService {
	return &sservice{r}
}

func (sr sservice) SetSession(s *entity.Session) error {
	s.Session = sr.GenerateSession()
	return sr.sessionRepo.Save(s)
}

func (sr sservice) GetUser(s string) (*entity.Session, error) {
	return sr.sessionRepo.GetOne(s)
}

func (sr sservice) RemoveSession(s *entity.Session) error {
	return sr.sessionRepo.Delete(s.Session)
}

func (sr sservice) GenerateSession() string {
	return uuid.New().String()
}
