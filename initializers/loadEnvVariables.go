package initializers

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnvVariables() {

	// Loads All The Enviornment Variables
	err := godotenv.Load()

	if err != nil {
		logrus.Error("Error Loading ENV Variables")
	}

	logrus.Info("Enviornment Varirables Loaded")
}
