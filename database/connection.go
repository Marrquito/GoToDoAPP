package database

import (
	"log"
	"os"
	"time"

	myLog "github.com/Marrquito/GoToDoAPP/common/logger"
	"github.com/Marrquito/GoToDoAPP/config"
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

	logger := myLog.NewLogger(config.ErrorLevel, "Db Connection")

	if err != nil {
		panic("Could not connect to the database")
	}

	logger.Info("DB Connected!")

	return db
}

var Database *gorm.DB = Connect()
