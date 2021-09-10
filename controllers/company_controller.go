package controllers

import (
	"strconv"

	dbMySql "github.com/boazsoftwares/Boaz.Api.Go/database"
	"github.com/boazsoftwares/Boaz.Api.Go/models"
	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID not integer",
		})
		return
	}

	db := dbMySql.GetDatabase()

	var empresa models.Company
	err = db.First(empresa, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "It was not possible find company: " + err.Error(),
		})
		return
	}

	c.JSON(200, empresa)
}

func AddCompany(c *gin.Context) {
	db := dbMySql.GetDatabase()

	var empresa models.Company

	err := c.ShouldBindJSON(empresa)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error JSON convert: " + err.Error(),
		})
		return
	}

	err = db.Create(empresa).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Erro register company" + err.Error(),
		})
		return
	}

	c.JSON(200, empresa)
}

func GetCompanies(c *gin.Context) {
	db := dbMySql.GetDatabase()
	var empresas []models.Company
	err := db.Find(empresas).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Error find company" + err.Error(),
		})

		return
	}

	c.JSON(200, empresas)
}
