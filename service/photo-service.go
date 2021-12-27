package service

import (
	"clean-architecture/entity"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type pservice struct {
	photoRepo entity.PhotoRepository
}

func NewPhotoService(r entity.PhotoRepository) entity.PhotoService {
	return &pservice{r}
}

func (pservice) HashCode(s string) string {
	h := hmac.New(sha256.New, []byte("arjun"))
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (ps pservice) CreatePhoto(r *http.Request, u *entity.User) (*[]entity.Photo, error) {
	var phs []entity.Photo
	urls, err := ps.UploadFiles(r)
	for _, url := range urls {
		p := entity.Photo{
			Url:  url,
			User: &entity.User{},
		}
		ps.photoRepo.Save(&p)
		phs = append(phs, p)
	}
	return &phs, err
}

func (ps pservice) UploadFiles(r *http.Request) ([]string, error) {
	var err error
	var urls []string
	for _, headers := range r.MultipartForm.File {
		// headers is of type array index,value
		for _, hdr := range headers {
			file, err := hdr.Open()
			if err != nil {
				return urls, err
			}
			defer file.Close()

			dir, err := os.Getwd()
			if err != nil {
				return urls, err
			}

			fa := strings.Split(hdr.Filename, ".")

			nf := ps.HashCode(hdr.Filename) + "." + fa[len(fa)-1]

			nfile, err := os.Create(filepath.Join(dir, "public", nf))
			if err != nil {
				return urls, err
			}

			_, err = io.Copy(nfile, file)
			if err != nil {
				return urls, err
			}

			url := os.Getenv("IMG_URL") + nf

			urls = append(urls, url)

			if err != nil {
				return urls, err
			}
		}
	}
	return urls, err
}

func (pservice) RemoveFiles(url string) error {
	if len(url) == 0 {
		return errors.New("internal server error")
	}
	ar := strings.Split(url, "/")
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Remove(filepath.Join(dir, "public", ar[len(ar)-1]))
	return err
}

func (ps pservice) DeletePhoto(id int64, uId int64) (*entity.Photo, error) {
	p, err := ps.FindPhoto(id)
	if err != nil {
		return p, err
	}
	if uId != p.User.Id {
		return p, errors.New("access denied")
	}
	err = ps.RemoveFiles(p.Url)
	if err != nil {
		return p, err
	}
	err = ps.photoRepo.Delete(id)
	return p, err
}

func (ps pservice) FindPhoto(id int64) (*entity.Photo, error) {
	return ps.photoRepo.GetOne(id)
}

func (ps pservice) FindPhotos() (*[]entity.Photo, error) {
	return ps.photoRepo.GetAll()
}
