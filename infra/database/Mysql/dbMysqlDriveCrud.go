package dbMysqlDrive

import (
	"fmt"

	entities "github.com/boazsoftwares/Boaz.Api.Go/domain/entities"
	databaseInterface "github.com/boazsoftwares/Boaz.Api.Go/infra/database/interface"
)

type crudGeneric struct {
	Crud databaseInterface.CrudGenericInterface
}

var Crud crudGeneric

func (crud crudGeneric) Find(newid int, err error) entities.Company {
	db := GetDatabase()

	var companyResult = entities.Company{}
	err = db.First(companyResult, newid).Error
	if err != nil {
		fmt.Println("Error: It was not possible find company: " + err.Error())
		return companyResult
	}
	return companyResult
}
