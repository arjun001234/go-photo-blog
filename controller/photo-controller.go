package controller

import (
	md "clean-architecture/controller/middleware"
	"clean-architecture/entity"
	"encoding/json"
	"net/http"
	"strconv"

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

	ps, err := pc.photoService.CreatePhoto(r, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (pc photoController) GetPhoto(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	pId := pr.ByName("id")

	if len(pId) == 0 {
		http.Error(w, "Photo id not provided", http.StatusBadRequest)
		return
	}

	toInt, _ := strconv.Atoi(pId)

	p, err := pc.photoService.FindPhoto(int64(toInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (pc photoController) DeletePhoto(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {

	pId := pr.ByName("id")

	if len(pId) == 0 {
		http.Error(w, "Photo id not provided", http.StatusBadRequest)
	}

	u := r.Context().Value(md.USER).(entity.User)

	toInt, _ := strconv.Atoi(pId)

	p, err := pc.photoService.DeletePhoto(int64(toInt), u.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
func (pc photoController) GetPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ps, err := pc.photoService.FindPhotos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
