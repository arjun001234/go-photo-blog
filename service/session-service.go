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

func (sr sservice) SetSession(s *entity.Session) (*entity.Session, error) {
	id := uuid.New().String()
	s.Session = id
	err := sr.sessionRepo.Save(s)
	return s, err
}

func (sr sservice) GetUser(s *entity.Session) error {
	err := sr.sessionRepo.GetOne(s)
	return err
}

func (sr sservice) RemoveSession(s *entity.Session) error {
	return sr.sessionRepo.Remove(s)
}
