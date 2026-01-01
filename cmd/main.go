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

	r.Use(util.LoggingMiddleware)
	rApi := r.PathPrefix("/api").Subrouter()

	routeAuth := rApi.PathPrefix("/users").Subrouter()
	routeAuth.Use(util.AuthorizationMiddleware)
	routeAuth.HandleFunc("/", hUser.Delete).Methods(http.MethodDelete)                // soft delete
	routeAuth.HandleFunc("/profile", hUser.GetProfile).Methods(http.MethodGet)        // get profile
	routeAuth.HandleFunc("/profile", hUser.UpdateProfile).Methods(http.MethodPatch)   // update proifle
	routeAuth.HandleFunc("/password", hUser.ChangePassword).Methods(http.MethodPatch) // change password

	rApi.HandleFunc("/users/register", hUser.Register).Methods(http.MethodPost)
	rApi.HandleFunc("/users/login", hUser.Login).Methods(http.MethodPost)
	// rApi.HandleFunc("/users/password/reset", hUser.ResetPassword).Methods(http.MethodGet) // recovery password

	// rApi.HandleFunc("/users/:id", hUser.Get).Methods(http.MethodGet)
	// rApi.HandleFunc("/users", hUser.List).Methods(http.MethodGet) // manager

	fmt.Println("Server Listening at *:7010")
	http.ListenAndServe(":7010", r)
}
