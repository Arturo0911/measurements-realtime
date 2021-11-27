package connection

// Implementing Factory dessign pattern
// for MySQL connections

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database credential is a configuration struct
// for a database connection
type DatabaseCredential struct {
	DSN          string
	DriverName   string
	DB           sql.DB
	DatabaseName string
	DBAddress    string
	//Connector  func() (*sql.DB, error)
}

var MysqlDB = &DatabaseCredential{}

func NewInstance() *DatabaseCredential {
	return &DatabaseCredential{
		DSN:          "",
		DriverName:   "mysql",
		DatabaseName: "GreenHouse",
	}
}

func (db *DatabaseCredential) DatabaseConnection() {
	dbConn, dbErr := sql.Open("mysql", fmt.Sprintf("root_payload:@tcp(127.0.0.1:3306)/%s", db.DatabaseName))
	if dbErr != nil {
		log.Fatalf("Error in the connection: %s", dbErr)
		return
	}
	MysqlDB.DB = *dbConn
	defer dbConn.Close()
}
