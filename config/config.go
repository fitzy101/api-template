package config

import (
	"os"
)

func DB_HOST() string {
	return os.Getenv("DB_HOST")
}

func DB_PASSWORD() string {
	return os.Getenv("DB_PASSWORD")
}

func DB_SCHEMA() string {
	return os.Getenv("DB_SCHEMA")
}

func DB_PORT() string {
	return os.Getenv("DB_PORT")
}

func DB_USERNAME() string {
	return os.Getenv("DB_USERNAME")
}
