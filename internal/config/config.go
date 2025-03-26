package config

import (
	"os"
	"strconv"
)

const (
	Port   = 7540
	Secret = "sdsjhUHAHu78sakkkj7878"
)

type Config struct {
	DBFile   string
	Port     string
	Password string
	Secret   string
}

func LoadConfig() *Config {
	port := strconv.Itoa(Port)
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		port = envPort
	}

	pass := ""
	envPass := os.Getenv("TODO_PASSWORD")
	if len(envPass) > 0 {
		pass = envPass
	}

	secret := Secret
	envSecret := os.Getenv("TODO_SECRET")
	if len(envSecret) > 0 {
		secret = envSecret
	}

	return &Config{
		DBFile:   os.Getenv("TODO_DBFILE"),
		Port:     port,
		Password: pass,
		Secret:   secret,
	}
}
