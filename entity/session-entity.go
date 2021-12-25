package entity

import "time"

type Session struct {
	Session   string    `json:"session"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type SessionRepository interface {
	Save(s *Session) error
	Remove(s *Session) error
	GetOne(s *Session) error
}

type SessionService interface {
	SetSession(s *Session) (*Session, error)
	GetUser(s *Session) error
	RemoveSession(s *Session) error
}
