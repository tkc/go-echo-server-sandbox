package env

import (
	"os"
	"github.com/joho/godotenv"
)

const EnvFile = ".env"

func IsProd() bool {
	if (GetAppEnv() == "prod") {
		return true
	} else {
		return false
	}
}

func GetAppEnv() string {
	godotenv.Load(EnvFile)
	return os.Getenv("APP_ENV")
}

func GetPort() string {
	godotenv.Load(EnvFile)
	return os.Getenv("PORT")
}

func GetDataBaseAccess() string {
	godotenv.Load(EnvFile)
	return os.Getenv("DATA_BASE")
}
