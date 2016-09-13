package config

import (
	"os"
)

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func DBSchema() string {
	return os.Getenv("DB_SCHEMA")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBUsername() string {
	return os.Getenv("DB_USERNAME")
}
