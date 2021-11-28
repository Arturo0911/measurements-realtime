package connection

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// Database credential is a configuration struct
// for a database connection
type DatabaseCredential struct {
	DSN string
	DB  gorm.DB
}

func NewInstance() *gorm.DB {
	return GetConnection()
}

func GetConnection() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open("root_payload:@tcp(127.0.0.1:3306)/GreenHouse?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Fatalf("Error by: %s", err)
		return nil
	}
	return db
}
