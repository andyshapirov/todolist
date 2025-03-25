package config

import (
	"os"
	"strconv"
)

const (
	PORT     = 7540
	PASSWORD = "12345"
	SECRET   = "sdsjhUHAHu78sakkkj7878"
	TOKEN    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZF9oYXNoIjoiNTk5NDQ3MWFiYjAxMTEyYWZjYzE4MTU5ZjZjYzc0YjRmNTExYjk5ODA2ZGE1OWIzY2FmNWE5YzE3M2NhY2ZjNSJ9.uAob9jXw9Nky_d6jcYyx964J5hLqrEtfm8TWU6HTDBY"
)

type Config struct {
	DBFile   string
	Port     string
	Password string
	Secret   string
}

func LoadConfig() *Config {
	port := strconv.Itoa(PORT)
	envPort := os.Getenv("TODO_PORT")
	if len(envPort) > 0 {
		port = envPort
	}

	pass := PASSWORD
	envPass := os.Getenv("TODO_PASSWORD")
	if len(envPass) > 0 {
		pass = envPass
	}

	secret := SECRET
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
