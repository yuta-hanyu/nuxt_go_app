package config

import (
	// "goland/utils"
	"log"
	"os"
	"github.com/joho/godotenv"
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
	// utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Config = ConfigList {
		DbUserName: os.Getenv("DB_USER"),
		DbPassWord: os.Getenv("DB_PASSWORD"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		LogFile: os.Getenv("ERROR_LOG"),

		// Port: cfg.Section("web").Key("port").MustString("8080"),
		// SQLDriver: cfg.Section("db").Key("driver").String(),
		// DB_NAMA: cfg.Section("db").Key("name").String(),
		// LogFile: cfg.Section("web").Key("logfile").String(),
		// Static: cfg.Section("web").Key("static").String(),
	}
}