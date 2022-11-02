package gorm_demo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./gorm_demo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
