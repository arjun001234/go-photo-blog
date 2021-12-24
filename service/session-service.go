package service

import "clean-architecture/entity"

type sservice struct {
	sessionRepo entity.SessionRepository
}

func NewSessionService(r entity.SessionRepository) *sservice {
	return &sservice{r}
}
