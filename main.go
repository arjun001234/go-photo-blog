package main

import (
	"clean-architecture/config"
	"clean-architecture/controller"
	"clean-architecture/controller/middleware"
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
	config.LoadTemplates()
}

func main() {
	db := storage.NewMySQLDatabase()
	defer db.Close()
	ur := repo.NewUserRepository(db)
	pr := repo.NewphotoRepository(db)
	sr := repo.NewsessionRepository(db)
	ps := service.NewPhotoService(pr)
	ss := service.NewSessionService(sr)
	us := service.NewUserService(ur, ss)
	uc := controller.NewUserController(us)
	pc := controller.NewPhotoController(ps)
	router := httprouter.New()
	md := middleware.NewMiddleware(us)
	router.ServeFiles("/public/*filepath", http.Dir("./public"))
	router.POST("/user", uc.NewUser)
	router.GET("/users", uc.GetUsers)
	router.POST("/user/login", uc.UserLogin)
	router.POST("/user/logout", middleware.Adapt(uc.UserLogout, md.Auth()))
	router.DELETE("/user", middleware.Adapt(uc.RemoveUser, md.Auth()))
	router.PATCH("/user", middleware.Adapt(uc.UpdateUser, md.Auth()))
	router.GET("/user", middleware.Adapt(uc.GetUser, md.Auth()))
	router.GET("/", pc.GetPhotos)
	router.GET("/photo/:id", pc.GetPhoto)
	router.DELETE("/photo/:id", middleware.Adapt(pc.DeletePhoto, md.Auth()))
	router.POST("/photo", middleware.Adapt(pc.NewPhoto, md.Auth()))
	err := http.ListenAndServe(`:`+os.Getenv("PORT"), router)
	if err != nil {
		fmt.Println("Server is up and running on port " + os.Getenv("PORT"))
	} else {
		log.Fatal(err)
	}
}
