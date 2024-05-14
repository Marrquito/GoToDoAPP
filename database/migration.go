package database

import (
	"fmt"

	"github.com/Marrquito/GoToDoAPP/common/colors"
	"github.com/Marrquito/GoToDoAPP/database/models"

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
	fmt.Println("🔧", colors.Bold("Migrating database\n"))
	fmt.Printf("  ◽ Running all migrations\n\n")

	db.AutoMigrate(
		&models.ToDo{},
	)

	__db, _ := db.DB()

	fmt.Printf("\n✅ %s\n", colors.Bold("Database migrated successfuly!"))
	__db.Close()
}
