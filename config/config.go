package config

import "fmt"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "eforl",
			Password: "password",
			Name:     "go_todo",
			Charset:  "utf8",
		},
	}
}
