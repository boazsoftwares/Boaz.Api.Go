package databaseInterface

type CrudGenericInterface interface {
	Add() bool
	Remove() bool
	Find() interface{}
	FindAll() []interface{}
	Update() bool
	Insert() bool
}
