package databaseInterface

type CrudGenericInterface interface {
	getDatabase()
	Add() bool
	Remove() bool
	Find(entity interface{}, newid int, err error) interface{}
	FindAll() []interface{}
	Update() bool
	Insert() bool
}
