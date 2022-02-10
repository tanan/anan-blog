package handler


type SQLHandler interface {
	Close() error
}

type ContoroDBHandler interface {
	SQLHandler
	FindByID(id string) (contract string, err error)
}