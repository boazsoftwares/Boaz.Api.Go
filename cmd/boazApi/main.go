package main

import (
	"fmt"

	"github.com/boazsoftwares/Boaz.Api.Go/api/server"
	dbMysqlDrive "github.com/boazsoftwares/Boaz.Api.Go/infra/database/Mysql"
)

func main() {
	fmt.Println("Start application GoLang")

	fmt.Println("Start database architecture")
	dbMysqlDrive.MysqlTest()

	fmt.Println("Start new server Company")
	server := server.NewServerCompany()
	server.RunServerCompany()
}
