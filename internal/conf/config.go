package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Path   string `yaml:"path"`
	Target string `yaml:"target"`
}

type Authenticate struct {
	Jwtsecret string `yaml:"jwtsecret"`
}

type Config struct {
	Authenticate Authenticate `yaml:"authenticate"`
	Routes       []Route      `yaml:"routes"`
}

func LoadConfig() {
	data, err := os.ReadFile("config.yaml") // load config file
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
}
