package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type Config struct {
	HealthPlanetDB  DBConfig        `yaml:"health_planet_database"`
	HealthPlanetAPI HealthPlanetAPI `yaml:"health_planet_api"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Schema   string `yaml:"schema"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
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

func (config DBConfig) GetConnInfo() string {
	return config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Schema + "?timeout=" + strconv.Itoa(TIMEOUT) + "s" + "&parseTime=true"
}

func LoadFile(path string) (*Config, error) {
	var config Config
	b, _ := os.ReadFile(path)
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) LoadEnvPassword() {
	p := os.Getenv("MYSQL_PASSWORD")
	if p != "" {
		c.HealthPlanetDB.Password = p
	}
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
