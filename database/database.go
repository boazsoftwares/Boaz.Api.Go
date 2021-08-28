package database

import (
	"log"
	"time"

	"github.com/boazsoftwares/Boaz.Api.Go/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectionConfig() {
	connectionString := "root:root@tcp(127.0.0.1:3306)/BoazApiGo?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigration(db)
}

func GetDatabase() *gorm.DB {
	return db
}
