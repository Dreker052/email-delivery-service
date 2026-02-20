package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SMTPHost string
	SMTPPort int
	SMTPUser string
	SMTPPass string

	RedisAddr string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err

	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: port,
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPass: os.Getenv("SMTP_PASS"),

		RedisAddr: os.Getenv("REDIS_ADDR"),
	}, nil
}
