package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT            string
	DATABASE_HOST   string
	DATABASE_PORT   string
	DATABASE_USER   string
	DATABASE_PASS   string
	DATABASE_DBNAME string
	BROKER_URL      string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		DATABASE_HOST = os.Getenv("DATABASE_HOST")
		DATABASE_PORT = os.Getenv("DATABASE_PORT")
		DATABASE_USER = os.Getenv("DATABASE_USER")
		DATABASE_PASS = os.Getenv("DATABASE_PASS")
		DATABASE_DBNAME = os.Getenv("DATABASE_DBNAME")
		BROKER_URL = os.Getenv("BROKER_URL")
	} else {
		PORT = viper.GetString("PORT")
		DATABASE_HOST = viper.GetString("DATABASE_HOST")
		DATABASE_PORT = viper.GetString("DATABASE_PORT")
		DATABASE_USER = viper.GetString("DATABASE_USER")
		DATABASE_PASS = viper.GetString("DATABASE_PASS")
		DATABASE_DBNAME = viper.GetString("DATABASE_DBNAME")
		BROKER_URL = viper.GetString("BROKER_URL")
	}

	if len(PORT) == 0 ||
		len(DATABASE_HOST) == 0 ||
		len(DATABASE_PORT) == 0 ||
		len(DATABASE_USER) == 0 ||
		len(DATABASE_PASS) == 0 ||
		len(DATABASE_DBNAME) == 0 ||
		len(BROKER_URL) == 0 {
		panic("Missing environment variables")
	}
}
