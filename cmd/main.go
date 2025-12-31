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
	"github.com/ilmsg/fluffy-octo-tribble/util"
)

func main() {

	db := database.NewDatabaseWithSqlite("fluffy-octo-tribble.db")
	db.AutoMigrate(&model.Project{}, &model.Task{}, &model.User{}, &model.Auth{}, &model.Profile{})

	rUser := repository.NewUserRepository(db)
	sUser := service.NewUserService(rUser)
	hUser := handler.NewUserWebHandler(sUser)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	routeAuth := r.PathPrefix("/users").Subrouter()
	routeAuth.Use(util.AuthorizationMiddleware)
	routeAuth.HandleFunc("/", hUser.Get).Methods(http.MethodGet)                     // get profile
	routeAuth.HandleFunc("/", hUser.Delete).Methods(http.MethodDelete)               // soft delete
	routeAuth.HandleFunc("/profile", hUser.UpdateProfile).Methods(http.MethodPatch)  // update proifle
	routeAuth.HandleFunc("/password", hUser.ChangePassword).Methods(http.MethodPost) // change password

	r.HandleFunc("/users/password/reset", hUser.ResetPassword).Methods(http.MethodGet) // recovery password
	r.HandleFunc("/users/register", hUser.Register).Methods(http.MethodPost)
	r.HandleFunc("/users/login", hUser.Login).Methods(http.MethodPost)

	r.HandleFunc("/users/:id", hUser.Get).Methods(http.MethodGet)
	r.HandleFunc("/users", hUser.List).Methods(http.MethodGet) // manager

	fmt.Println("Server Listening at *:7010")
	http.ListenAndServe(":7010", r)
}
