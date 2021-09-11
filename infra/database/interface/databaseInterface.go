package databaseInterface

type DatabaseInterface interface {
	ConnectionConfig()
	OpenConnection()
	ClosedConnection()
	GetDatabase()
	Migration()
}
