package config

import "github.com/bruno-sca/tasqcoin-back-go/pkg/env"

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Server struct {
		Port string
	}
}

func NewConfig() *Config {
	return &Config{
		Database: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		}{
			Host:     env.GetEnvOrDie("DB_HOST"),
			Port:     env.GetEnvOrDie("DB_PORT"),
			User:     env.GetEnvOrDie("DB_USER"),
			Password: env.GetEnvOrDie("DB_PASSWORD"),
			Name:     env.GetEnvOrDie("DB_NAME"),
		},

		Server: struct {
			Port string
		}{
			Port: env.GetEnvOrDie("SERVER_PORT"),
		},
	}
}
