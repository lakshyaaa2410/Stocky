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

	dsn := "postgresql://neondb_owner:npg_QEpA8rVvUI7N@ep-shy-lab-a139e4ge-pooler.ap-southeast-1.aws.neon.tech/assignment?sslmode=require&channel_binding=require"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Error Connecting To Database")
		return
	}

	fmt.Println("Connected To Database!")
}
