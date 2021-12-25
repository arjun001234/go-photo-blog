package controller

import (
	md "clean-architecture/controller/middleware"
	"clean-architecture/entity"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type photoController struct {
	photoService entity.PhotoService
}

func NewPhotoController(s entity.PhotoService) *photoController {
	return &photoController{s}
}

func (pc photoController) NewPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := r.Context().Value(md.USER).(entity.User)

	r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Images size too large", http.StatusInternalServerError)
		return
	}

	p := entity.Photo{
		User: &u,
	}

	err := pc.photoService.CreatePhoto(r, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	w.WriteHeader(201)
}

func (photoController) GetPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Existing Photo"))
}

func (photoController) DeletePhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Delete Photo"))
}
func (photoController) GetPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Existing Photos"))
}
