package entity

import "time"

type Photo struct {
	Id        int64     `json:"id,omitempty"`
	Url       string    `json:"url"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoRepository interface {
	Save(p *Photo) error
	Delete(id int64) error
	Update(id int64) error
	GetOne(id int64) (*Photo, error)
	GetAll() (*[]Photo, error)
}

type PhotoService interface {
	CreatePhoto(p *Photo)
	UpdatePhoto(p *Photo)
	DeletePhoto(id int64)
	FindPhoto(id int64)
	FindPhotos()
}
