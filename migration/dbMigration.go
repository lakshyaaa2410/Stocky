package main

import (
	"fmt"

	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	err := initializers.DB.AutoMigrate(&models.Reward{})

	if err != nil {
		logrus.Error("Error Migrating Models")
		return
	}

	fmt.Println("Migration Done")
}
