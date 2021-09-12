package dbMysqlDrive

import (
	"fmt"
	"reflect"

	databaseInterface "github.com/boazsoftwares/Boaz.Api.Go/infra/database/interface"
	"gorm.io/gorm"
)

type crudGeneric struct {
	Crud databaseInterface.CrudGenericInterface
}

var db *gorm.DB

func Crud() crudGeneric {
	db = GetDatabase()
	return crudGeneric{}
}

func (crud crudGeneric) Find(entity interface{}, newid int, err error) interface{} {
	err = db.First(entity, newid).Error
	if err != nil {
		fmt.Println("Error: It was not possible find " + reflect.TypeOf(entity).String() + ": " + err.Error())
		return entity
	}
	return entity
}
