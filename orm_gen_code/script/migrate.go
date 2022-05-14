package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dsn = "./gen_sqlite.db"

type User struct {
	UUID    string `gorm:"primary_key"`
	Name    string `gorm:"type:varchar(255);not null"`
	Age     int    `gorm:"type:int;not null"`
	Version int    `gorm:"type:int;not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
