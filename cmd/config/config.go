package config

import (
	"fmt"
	"os"
)

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
	SecretKey string
}

func New() *Config {
	cfg := &Config{}

	cfg.Database.Host = os.Getenv("DATABASE_HOST")
	cfg.Database.Port = os.Getenv("DATABASE_PORT")
	cfg.Database.User = os.Getenv("DATABASE_USER")
	cfg.Database.Password = os.Getenv("DATABASE_PASSWORD")
	cfg.Database.Name = os.Getenv("DATABASE_NAME")
	cfg.SecretKey = os.Getenv("SECRET_KEY")
	cfg.Server.Port = ":" + os.Getenv("SERVER_PORT")

	return cfg
}

func (c *Config) GetDB() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)
}
