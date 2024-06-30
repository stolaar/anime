package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     int
}

func LoadDBConfig() DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
	}
}
