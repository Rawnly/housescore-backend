package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"housescore/configuration"
	"log"
	"strings"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
 	dbConfig := configuration.Config.Database

 	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

 	if err != nil {
 		panic(err)
	}

	if strings.ToLower(configuration.Config.Log.Level) == "debug" {
		log.Printf("Connecting to: %s \n", dsn)
	}

	return db
}


func Instance() *gorm.DB {
	if db == nil {
		return Init()
	}

	return db
}


