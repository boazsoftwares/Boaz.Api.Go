package dbMysqlDrive

import (
	"log"
	"time"

	databaseInterface "github.com/boazsoftwares/Boaz.Api.Go/infra/database/interface"
	"github.com/boazsoftwares/Boaz.Api.Go/infra/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDatabase *gorm.DB
var driverStart DBMySqlConfig

type ConnectionConfig struct {
	Host     string
	Port     int
	Database string
	Password string
}

type DBMySqlConfig struct {
	databaseInterface.DatabaseInterface
	Config ConnectionConfig
	databaseInterface.CrudGenericInterface
}

func (drive DBMySqlConfig) Migration() {
	migrations.RunMigration(gormDatabase)
}

func GetDatabase() *gorm.DB {
	return gormDatabase
}

func (drive DBMySqlConfig) GetDatabase() *gorm.DB {
	return gormDatabase
}

func (drive DBMySqlConfig) OpenConnection() {
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

func (drive DBMySqlConfig) ClosedConnection() {
	config, _ := gormDatabase.DB()
	config.Close()
}

func MysqlTest() DBMySqlConfig {
	driverStart.OpenConnection()
	driverStart.Migration()
	driverStart.ClosedConnection()
	return driverStart
}
