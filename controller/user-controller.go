package controller

import (
	md "clean-architecture/controller/middleware"
	"clean-architecture/entity"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type userController struct {
	userService entity.UserService
}

func NewUserController(s entity.UserService) *userController {
	return &userController{s}
}

func (uc userController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	us, err := uc.userService.FindUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(us)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (uc userController) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := r.Context().Value(md.USER).(entity.User)

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (uc userController) NewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var u entity.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := uc.userService.CreateUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    s.Session,
		HttpOnly: true,
	})

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (uc userController) UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var u entity.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := uc.userService.ValidateCredential(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    s.Session,
		HttpOnly: true,
	})

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (uc userController) UserLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := r.Context().Value(md.USER).(entity.User)

	ck, _ := r.Cookie("session")

	ck.MaxAge = -1

	http.SetCookie(w, ck)

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (uc userController) RemoveUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := r.Context().Value(md.USER).(entity.User)

	err := uc.userService.DeleteUser(u.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}

func (uc userController) UpdateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := r.Context().Value(md.USER).(entity.User)

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uc.userService.UpdateUser(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
