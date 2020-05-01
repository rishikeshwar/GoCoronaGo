package controllers

import (
	"log"
	"myapp/app/models"

	//idk
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

//DB is the variable handing database operations
var DB *gorm.DB

//InitDB Initialises the Ddatabase and helps to create new Tables
func InitDB() {

	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	println(dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Test{})
	DB = db
}
