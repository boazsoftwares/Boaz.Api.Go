package controllers

import (
	"strconv"

	"github.com/boazsoftwares/Boaz.Api.Go/database"
	"github.com/boazsoftwares/Boaz.Api.Go/models"
	"github.com/gin-gonic/gin"
)

func GetEmpresa(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID não é inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var empresa models.Empresa
	err = db.First(empresa, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Não foi possível consultar a empresa: " + err.Error(),
		})
		return
	}

	c.JSON(200, empresa)
}

func AddEmpresa(c *gin.Context) {
	db := database.GetDatabase()

	var empresa models.Empresa

	err := c.ShouldBindJSON(empresa)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Erro na conversão do JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(empresa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Erro no cadastro da empresa" + err.Error(),
		})
		return
	}

	c.JSON(200, empresa)
}

func GetEmpresas(c *gin.Context) {
	db := database.GetDatabase()
	var empresas []models.Empresa
	err := db.Find(empresas).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Erro ao consultar empresas" + err.Error(),
		})

		return
	}

	c.JSON(200, empresas)
}
