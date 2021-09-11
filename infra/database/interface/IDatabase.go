package iDatabase

import "fmt"

type ConnectionConfig struct {
	Host     string
	Port     int
	Database string
	Password string
}

type DatabaseBase interface {
	ConnectionConfig()
	OpenConnection()
	ClosedConnection()
	GetDatabase()
}

type DatabaseDriver struct {
	DatabaseBase
	Config ConnectionConfig
}

func (sa DatabaseDriver) ClosedConnection() {
	fmt.Println(sa)
}

func IDatebaseExecutTest() DatabaseDriver {
	var driver DatabaseDriver
	driver.ClosedConnection()
	return driver
}
