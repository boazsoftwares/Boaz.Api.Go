package main

import (
	"fmt"

	"github.com/boazsoftwares/Boaz.Api.Go/server"
)

func main() {
	fmt.Println("Start server BoazApiGo")
	server := server.NewServer()

	server.Run()
}
