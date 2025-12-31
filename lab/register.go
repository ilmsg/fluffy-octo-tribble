//go:build ignore

package main

import (
	"fmt"
	"log"

	"github.com/ilmsg/fluffy-octo-tribble/database"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"github.com/ilmsg/fluffy-octo-tribble/repository"
)

func main() {
	db := database.NewDatabaseWithSqlite("fluffy-octo-tribble.db")
	db.Migrator().DropTable(&model.Project{}, &model.Task{}, &model.User{}, &model.Auth{}, &model.Profile{})
	db.AutoMigrate(&model.Project{}, &model.Task{}, &model.User{}, &model.Auth{}, &model.Profile{})

	rUser := repository.NewUserRepository(db)
	// sUser := service.NewUserService(rUser)

	registerDto := model.RegisterDto{
		Email:           "scott@gmail.com",
		Password:        "tiger",
		ConfirmPassword: "tiger",
		Name:            "Scott Tiger",
		Location:        "Bangkok",
	}

	newUser, err := rUser.Register(&registerDto)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Auth: %+v\nProfile: %+v\n", newUser.Auth, newUser.Profile)

	// newUser, err := sUser.Register(&registerDto)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", newUser)

	// loginDto := model.LoginDto{Email: "scott@gmail.com", Password: "tiger"}
	// auth, err := sUser.Login(&loginDto)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", auth)
}
