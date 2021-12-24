package entity

import "time"

type Session struct {
	Session   string    `json:"session"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type SessionRepository interface {
	Save(s *Session) error
	Delete(id int64) error
	GetOne(id int64) (*Session, error)
}
