package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type databaseRule struct {
	db    *gorm.DB
	table string
}

func (r databaseRule) rewrite(path string) *string {
	return nil
}

func newDatabaseRule() *databaseRule {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return &(databaseRule{
		db: db,
	})
}
