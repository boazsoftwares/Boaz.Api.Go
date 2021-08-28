package main

import (
	"fmt"

	"github.com/boazsoftwares/Boaz.Api.Go/database"
	"github.com/boazsoftwares/Boaz.Api.Go/server"
)

func main() {
	fmt.Println("Start Database")
	database.ConnectionConfig()

	fmt.Println("Start Server BoazApiGo")
	server := server.NewServer()

	server.Run()
}
