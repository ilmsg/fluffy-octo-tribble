package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilmsg/fluffy-octo-tribble/database"
	"github.com/ilmsg/fluffy-octo-tribble/handler"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"github.com/ilmsg/fluffy-octo-tribble/repository"
	"github.com/ilmsg/fluffy-octo-tribble/service"
)

func main() {

	db := database.NewDatabaseWithSqlite("fluffy-octo-tribble.db")
	db.AutoMigrate(&model.Project{}, &model.Task{}, &model.User{})

	rUser := repository.NewUserRepository(db)
	sUser := service.NewUserService(rUser)
	hUser := handler.NewUserWebHandler(sUser)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	// update proifle
	r.HandleFunc("/users/profile", hUser.UpdateProfile).Methods(http.MethodPatch)
	// change password
	r.HandleFunc("/users/password", hUser.ChangePassword).Methods(http.MethodPatch)

	r.HandleFunc("/users/:id", hUser.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/:id", hUser.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/users", hUser.Create).Methods(http.MethodPost)
	r.HandleFunc("/users", hUser.List).Methods(http.MethodGet)

	fmt.Println("Server Listening at *:7010")
	http.ListenAndServe(":7010", r)
}
