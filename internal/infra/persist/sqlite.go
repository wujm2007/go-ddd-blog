package persist

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
