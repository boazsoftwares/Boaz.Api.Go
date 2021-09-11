package databaseInterface

type CrudGenericInterface interface {
	getDatabase()
	Add() bool
	Remove() bool
	Find(newid int) interface{}
	FindAll() []interface{}
	Update() bool
	Insert() bool
}
