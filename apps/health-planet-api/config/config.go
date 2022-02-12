package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	HealthPlanetAPI HealthPlanetAPI `yaml:"health_planet_api"`
}

type HealthPlanetAPI struct {
	Id           string
	Secret       string
	AccessToken  string
	RefreshToken string
}

const (
	TIMEOUT = 5
)

func LoadFile(path string) (*Config, error) {
	var config Config
	b, _ := os.ReadFile(path)
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) LoadEnvHealthPlanet() {
	id := os.Getenv("HEALTH_PLANET_ID")
	if id != "" {
		c.HealthPlanetAPI.Id = id
	}
	refreshToken := os.Getenv("HEALTH_PLANET_REFRESH_TOKEN")
	if refreshToken != "" {
		c.HealthPlanetAPI.RefreshToken = refreshToken
	}
	secret := os.Getenv("HEALTH_PLANET_SECRET")
	if secret != "" {
		c.HealthPlanetAPI.Secret = secret
	}
}
