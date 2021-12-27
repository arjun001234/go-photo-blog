package entity

import "time"

type Session struct {
	Session   string    `json:"session"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type SessionRepository interface {
	Save(s *Session) error
	Delete(s string) error
	GetOne(s string) (*Session, error)
}

type SessionService interface {
	SetSession(s *Session) error
	GetUser(s string) (*Session, error)
	RemoveSession(s *Session) error
	GenerateSession() string
}
