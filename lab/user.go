//go:build ignore

package main

import (
	"log"

	"github.com/ilmsg/fluffy-octo-tribble/database"
	"github.com/ilmsg/fluffy-octo-tribble/model"
)

func main() {
	db := database.NewDatabaseWithSqlite("fluffy-octo-tribble.db")
	db.Migrator().DropTable(&model.Project{}, &model.Task{}, &model.User{}, &model.Auth{}, &model.Profile{})
	db.AutoMigrate(&model.Project{}, &model.Task{}, &model.User{}, &model.Auth{}, &model.Profile{})

	user := model.User{}
	if tx := db.Create(&user); tx.Error != nil {
		log.Fatal(tx.Error)
	}
	log.Printf("%+v", user)

	auth := model.Auth{Email: "scott@gmail.com", Password: "tiger", UserId: user.ID}
	if tx := db.Create(&auth); tx.Error != nil {
		log.Fatal(tx.Error)
	}
	log.Printf("%+v", auth)

	profile := model.Profile{Name: "Scott Tiger", Location: "Bangkok", UserId: user.ID}
	if tx := db.Create(&profile); tx.Error != nil {
		log.Fatal(tx.Error)
	}
	log.Printf("%+v", profile)

	var users []model.User
	if tx := db.Preload("Auth").Preload("Profile").Find(&users); tx.Error != nil {
		log.Fatal(tx.Error)
	}
	for _, user := range users {
		log.Printf("%+v", user)
	}
}
