package gdb

import "github.com/jinzhu/gorm"

var db *gorm.DB

func SetDB(gdb *gorm.DB) {
	db = gdb
}

func GetDB() *gorm.DB {
	return db
}
