package handler

import (
	"health-planet-api/domain"
	"net/http"
)

func (h Handler) Me(ctx Context) error {
	return ctx.JSON(http.StatusOK, domain.UserHealth{
		BirthDate: "BirthDate",
		Height:    "Height",
		Sex:       "Sex",
		Data: []domain.HealthData{
			{
				"test",
				"test",
				"test",
				"6000",
			},
		},
	})
}
