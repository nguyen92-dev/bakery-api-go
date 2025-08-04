package database

import "gorm.io/gorm"

type PreloadEntity struct {
	Entity string
}

func Preload(db *gorm.DB, preload []PreloadEntity) *gorm.DB {
	for _, item := range preload {
		db = db.Preload(item.Entity)
	}
	return db
}
