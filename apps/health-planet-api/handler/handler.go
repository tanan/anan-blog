package handler

import (
	"health-planet-api/infrastracture/database"
)

type Handler struct {
	healthPlanetDBHandler HealthPlanetDBHandler
}

func NewHandler(healthPlanetDBInfo string) (*Handler, error) {
	healthPlanetDBHandler, err := database.NewSQLHandler(healthPlanetDBInfo)
	if err != nil {
		return nil, err
	}
	return &Handler{
		healthPlanetDBHandler: healthPlanetDBHandler,
	}, nil
}
