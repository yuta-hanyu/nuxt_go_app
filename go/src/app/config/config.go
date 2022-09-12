package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"goland/utils"
)

type ConfigList struct {
	DbUserName string
	DbPassWord string
	DbHost string
	DbPort string
	DbName string
	LogFile string
	// Static string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(os.Getenv("ERROR_LOG"))
}

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Config = ConfigList {
		DbUserName: os.Getenv("DB_USER"),
		DbPassWord: os.Getenv("DB_PASSWORD"),
		DbHost: os.Getenv("DB_HOSTS"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		LogFile: os.Getenv("ERROR_LOG"),
	}
}