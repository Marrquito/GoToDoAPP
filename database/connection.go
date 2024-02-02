package database

import (
	"log"
	"os"
	"time"

	myLog "github.com/Marrquito/GoToDoAPP/common/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dns := "host=localhost user=todo-local password=local_local dbname=todo-local port=5438 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})

	logger := myLog.NewLogger(1, "Db Connection")

	if err != nil {
		logger.Error("Failed to connect to db: ", err)
	}

	logger.Info("DB Connected!")

	return db
}

var Database *gorm.DB = Connect()
