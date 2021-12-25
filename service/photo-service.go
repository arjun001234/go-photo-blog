package service

import (
	"clean-architecture/entity"
	"crypto/hmac"
	"crypto/sha256"
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

func (ps pservice) CreatePhoto(r *http.Request, p *entity.Photo) error {
	return ps.HandleFiles(r, p)
}

func (ps pservice) HandleFiles(r *http.Request, p *entity.Photo) error {
	var err error
	//    var fileUrl = "http://localhost:8080/public/"
	for _, headers := range r.MultipartForm.File {
		// headers is of type array index,value
		for _, hdr := range headers {
			file, err := hdr.Open()
			if err != nil {
				return err
			}
			defer file.Close()

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			fa := strings.Split(hdr.Filename, ".")

			nf := ps.HashCode(hdr.Filename) + "." + fa[len(fa)-1]

			nfile, err := os.Create(filepath.Join(dir, "public", nf))
			if err != nil {
				fmt.Println(err)
				return err
			}

			_, err = io.Copy(nfile, file)
			if err != nil {
				fmt.Println(err)
				return err
			}

			p.Url = os.Getenv("IMG_URL") + nf

			err = ps.photoRepo.Save(p)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return err
}

func (pservice) DeletePhoto(id int64) {}

func (pservice) FindPhoto(id int64) {}

func (pservice) FindPhotos() {}
