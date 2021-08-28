package routers

import (
	"github.com/boazsoftwares/Boaz.Api.Go/controllers"
	"github.com/gin-gonic/gin"
)

func ContigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		empresas := main.Group("empresas")
		{
			empresas.GET("/:id", controllers.GetEmpresa)
			empresas.GET("/", controllers.GetEmpresas)
			empresas.POST("/", controllers.AddEmpresa)
		}
	}

	return router
}
