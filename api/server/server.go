package server

import (
	"log"

	"github.com/boazsoftwares/Boaz.Api.Go/api/server/routers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServerCompany() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) RunServerCompany() {
	router := routers.ContigureRoutesCompany(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
