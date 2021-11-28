package connection

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

const DB_USERNAME = "root_payload"
const DB_PASSWORD = ""

func NewInstance() *gorm.DB {
	return GetConnection()
}

func GetConnection() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(DB_USERNAME+":"+DB_PASSWORD+"@tcp(127.0.0.1:3306)/GreenHouse?parseTime=true&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Fatalf("Error by: %s", err)
		return nil
	}
	return db
}
