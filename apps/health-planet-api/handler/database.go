package handler

type SQLHandler interface {
	Close() error
}

type HealthPlanetDBHandler interface {
	SQLHandler
	FindByID(id string) (contract string, err error)
}
