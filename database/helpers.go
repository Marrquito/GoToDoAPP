package database

import "gorm.io/gorm"

func LoadRelations(db *gorm.DB, relations *[]string) *gorm.DB {
	if relations == nil {
		return db
	}

	for _, relation := range *relations {
		db = db.Preload(relation)
	}

	return db
}
