package main

import (
	"fmt"

	"github.com/boazsoftwares/Boaz.Api.Go/api/server"
	dbMySql "github.com/boazsoftwares/Boaz.Api.Go/infra/database"
	iDatabase "github.com/boazsoftwares/Boaz.Api.Go/infra/database/interface"
)

func main() {
	fmt.Println("Start application GoLang")

	fmt.Println("Start database architecture")
	iDatabase.IDatebaseExecutTest()
	dbMySql.OpenConnection()

	fmt.Println("Start new server Company")
	server := server.NewServerCompany()
	server.RunServerCompany()
}
