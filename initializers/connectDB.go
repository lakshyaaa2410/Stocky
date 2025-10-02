package initializers

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Using Database Connection String, Hosted On Neon.
	DB_STRING := os.Getenv("DB_STRING")

	if DB_STRING == "" {
		logrus.Error("Database URL Is Not Defined, Please Define In ENV File")
		return
	}

	// Establishing A Connection.
	DB, err = gorm.Open(postgres.Open(DB_STRING), &gorm.Config{})

	if err != nil {
		logrus.Error("Error Connecting To Database")
		return
	}

	fmt.Println("Connected To Database!")
}
