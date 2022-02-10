package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type Config struct {
	ContoroDB DBConfig `yaml:"contoro_database"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Schema   string `yaml:"schema"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
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
		c.ContoroDB.Password = p
	}
}