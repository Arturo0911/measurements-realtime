package connection

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// const DB_USERNAME = "root"
// const DB_PASSWORD = "root"
// USER+":"+PASSWORD+"@tcp(127.0.0.1:3306)/GreenHouseRealTime?parseTime=true&loc=Local"

func NewInstance() *gorm.DB {
	return GetConnection()
}

func GetConnection() *gorm.DB {

	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD,
		HOST, DBNAME)

	db, err := gorm.Open(
		mysql.Open(URL),
		&gorm.Config{})
	if err != nil {
		log.Fatalf("Error by: %s", err)
		return nil
	}
	fmt.Println("Database connected successfully")
	return db
}
