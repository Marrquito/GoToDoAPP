package database

import (
	"TodoAppGO/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateAll(disableForeignKey bool) {
	dns := "host=localhost user=todo-local password=local_local dbname=todo-local port=5438 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: disableForeignKey,
	})

	if err != nil {
		panic("Error to connect to the database")
	}

	db.AutoMigrate(
		&models.ToDo{},
	)

	__db, _ := db.DB()
	__db.Close()
}
