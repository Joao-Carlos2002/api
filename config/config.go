package config

import (
	"github.com/Joao-Carlos2002/Api-User/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func initalizeDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

func Init() {
	Db = initalizeDb()
}
