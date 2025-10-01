package initializers

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=localhost user=postgres password=12345678 dbname=assignment port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Error Connecting To Database")
		return
	}

	fmt.Println("Connected To Database!")
}
