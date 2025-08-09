package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct{
	Host string `json:"host"`
	Port int `json:"port"`
}

type Config struct {
	Server Server `yaml:"server"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}
	return &cfg, nil
}

