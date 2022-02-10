package handler

import (
	"health-planet-api/infra/database"
)

type Handler struct {
	contoroDBHandler ContoroDBHandler
}

func NewHandler(contoroDBInfo string) (*Handler, error) {
	contoroDBHandler, err := database.NewSQLHandler(contoroDBInfo)
	if err != nil {
		return nil, err
	}
	return &Handler{
		contoroDBHandler: contoroDBHandler,
	}, nil
}