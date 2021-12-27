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
	// GetByUserAndPhotoId(p *Photo) error
}

type PhotoService interface {
	CreatePhoto(r *http.Request, u *User) (*[]Photo, error)
	DeletePhoto(id int64, uId int64) (*Photo, error)
	FindPhoto(id int64) (*Photo, error)
	FindPhotos() (*[]Photo, error)
	HashCode(s string) string
	UploadFiles(r *http.Request) ([]string, error)
	RemoveFiles(url string) error
}
