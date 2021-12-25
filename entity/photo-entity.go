package entity

import (
	"net/http"
	"time"
)

type Photo struct {
	Id        int64     `json:"id,omitempty"`
	Url       string    `json:"url"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoRepository interface {
	Save(p *Photo) error
	Delete(id int64) error
	GetOne(id int64) (*Photo, error)
	GetAll() (*[]Photo, error)
}

type PhotoService interface {
	CreatePhoto(r *http.Request, p *Photo) error
	DeletePhoto(id int64)
	FindPhoto(id int64)
	FindPhotos()
	HashCode(s string) string
	HandleFiles(r *http.Request, p *Photo) error
}
