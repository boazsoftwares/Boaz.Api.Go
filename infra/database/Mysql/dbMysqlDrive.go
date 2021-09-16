package dbMysqlDrive

import (
	"log"
	"time"

	"github.com/boazsoftwares/Boaz.Api.Go/infra/database"
	"github.com/boazsoftwares/Boaz.Api.Go/infra/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDatabase *gorm.DB
var databaseProvider DatabaseProviderConfig

type ConnectionConfig struct {
	Host     string
	Port     int
	Database string
	Password string
}

type DatabaseProviderConfig struct {
	database.Provider
	database.Crud
	Config ConnectionConfig
	Type   database.DatabaseTypeEnum
}

func (drive DatabaseProviderConfig) Migration() {
	migrations.RunMigration(gormDatabase)
}

func GetDatabase() *gorm.DB {
	return gormDatabase
}

func (driveConfig DatabaseProviderConfig) OpenConnection() {
	dbGorm, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/BoazApiGo?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	gormDatabase = dbGorm

	config, _ := gormDatabase.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func (drive DatabaseProviderConfig) ClosedConnection() {
	config, _ := gormDatabase.DB()
	config.Close()
}

func (drive DatabaseProviderConfig) GetDatabase() *gorm.DB {
	return gormDatabase
}

func MysqlTest() DatabaseProviderConfig {
	//databaseProvider.OpenConnection()
	//databaseProvider.Migration()
	//databaseProvider.ClosedConnection()
	return databaseProvider
}
