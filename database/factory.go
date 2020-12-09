package database

//********* TO DO ***********

// IFactory of database
type IFactory interface {
}

// DBTypes to be yieldded
type DBTypes string

// types available on DatabaseType
const (
	MongoDB DBTypes = "mongodb"
)

// Factory of databases
func Factory(db DBTypes) IFactory {
	switch db {
	case MongoDB:
		return newMongoConnection()
	default:
		return newMongoConnection()
	}
}
