package handler

import (
	"health-planet-api/config"
)

type Client interface {
	Get(url string, header map[string]string) (code int, respBody []byte, err error)
	Post(url string, header map[string]string, body []byte) (code int, respBody []byte, err error)
}

type Handler struct {
	APIConfig *config.HealthPlanetAPI
	Client    Client
}

func NewHandler(con *config.Config, client Client) (*Handler, error) {
	return &Handler{
		APIConfig: &con.HealthPlanetAPI,
		Client:    client,
	}, nil
}
