package main

import (
	"clean-architecture/config"
	"clean-architecture/controller"
	repo "clean-architecture/repo/mysql"
	"clean-architecture/service"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func init() {
	config.LoadConfig()
	db := storage.NewMySQLDatabase()
	ur := repo.NewUserRepository(db)
	pr := repo.NewphotoRepository(db)
	sr := repo.NewsessionRepository(db)
	us := service.NewUserService(ur)
	ps := service.NewPhotoService(pr)
	service.NewSessionService(sr)
	controller.NewUserController(us)
	controller.NewPhotoController(ps)
}

func main() {
	router := httprouter.New()
	err := http.ListenAndServe(`:`+os.Getenv("PORT"), router)
	if err != nil {
		fmt.Println("Server is up and running on port " + os.Getenv("PORT"))
	} else {
		log.Fatal(err)
	}
}
