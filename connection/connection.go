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
	DSN        string
	DriverName string
	Connector  func() (*sql.DB, error)
}

// DatabaseFactory contains all database
// credentials and instances
type DatabaseFactory struct {
	credentials map[string]DatabaseCredential
	instances   map[string]*DB
}

type DB struct {
	*sqlx.DB
}

type IDBConnection interface {
	Connect()
}

type MySQLConnection struct {
	connectionString string
}

func DataBaseConnection() {
	fmt.Println("[*] Starting the connection with database")

	db, err := sql.Open("mysql", "root_payload:@tcp(127.0.0.1:3306)/GreenHouse")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("[*] Connection Done!!")

	defer db.Close()
}
