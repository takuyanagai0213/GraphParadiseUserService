package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (database *gorm.DB) {
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
