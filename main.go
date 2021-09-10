package main

import (
	"fmt"

	dbMySql "github.com/boazsoftwares/Boaz.Api.Go/database"
	iDatabase "github.com/boazsoftwares/Boaz.Api.Go/database/interface"
	"github.com/boazsoftwares/Boaz.Api.Go/server"
)

func main() {
	iDatabase.IDatebaseExecutTest()
	fmt.Println("Start Application GoLang")
	dbMySql.OpenConnection()

	fmt.Println("Start Server BoazApiGo")
	server := server.NewServer()

	server.Run()
}
