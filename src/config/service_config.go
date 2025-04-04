package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type ServiceConfig struct {
	AppDataDirectory string
	ComicDirectory   string
}

func LoadServiceConfig() (*ServiceConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	appDataDirectory := os.Getenv("APP_DATA_DIRECTORY")
	if appDataDirectory == "" {
		appDataDirectory = "./data"
	}
	comicDirectory := os.Getenv("COMIC_DIRECTORY")
	if comicDirectory == "" {
		comicDirectory = "./comics"
	}
	return &ServiceConfig{
		AppDataDirectory: appDataDirectory,
		ComicDirectory:   comicDirectory,
	}, nil
}
