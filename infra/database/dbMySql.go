package dbMySql

import (
	"log"
	"time"

	"github.com/boazsoftwares/Boaz.Api.Go/infra/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlTeste(databaseGaneric ...interface{}) {

}

func GetConnectionString() string {
	return "root:root@tcp(127.0.0.1:3306)/BoazApiGo?charset=utf8mb4&parseTime=True&loc=Local"
}

var dataBaseTemp *gorm.DB

func OpenConnection() {
	database, err := gorm.Open(mysql.Open(GetConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	dataBaseTemp = database

	config, _ := dataBaseTemp.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func ExecuteMigration() {
	migrations.RunMigration(GetDatabase())
}

func GetDatabase() *gorm.DB {
	return dataBaseTemp
}
