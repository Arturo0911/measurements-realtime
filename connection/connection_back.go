package connection

// // Implementing Factory dessign pattern
// // for MySQL connections

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	. "github.com/Arturo0911/measurements-realtime/server"
// 	_ "github.com/go-sql-driver/mysql"
// )

// // Database credential is a configuration struct
// // for a database connection
// type DatabaseCredential struct {
// 	User         string
// 	Password     string
// 	DSN          string
// 	DriverName   string
// 	DB           sql.DB
// 	DatabaseName string
// 	DBAddress    string
// 	//Connector  func() (*sql.DB, error)
// }

// var MysqlDB = &DatabaseCredential{
// 	DSN:          "",
// 	DriverName:   "mysql",
// 	DatabaseName: "GreenHouse",
// }

// func NewInstance() *DatabaseCredential {
// 	return &DatabaseCredential{
// 		User:         "root_payload",
// 		Password:     "",
// 		DSN:          "",
// 		DriverName:   "mysql",
// 		DatabaseName: "GreenHouse",
// 		DBAddress:    "tcp(127.0.0.1:3306)",
// 	}
// }

// func (db *DatabaseCredential) DatabaseConnection() {
// 	dbConn, dbErr := sql.Open("mysql", fmt.Sprintf("root_payload:@%s/%s", db.DBAddress, db.DatabaseName))
// 	if dbErr != nil {
// 		log.Fatalf("Error in the connection: %s", dbErr)
// 		return
// 	}
// 	fmt.Println("Connection Done")
// 	db.DB = *dbConn
// 	defer dbConn.Close()
// }

// func (db *DatabaseCredential) GetValues() {
// 	var response []MeasurementsResponse
// 	results, err := db.DB.Query("SELECT temperaturaMin, temperaturaMax from niveles;")
// 	if err != nil {
// 		log.Fatalf("Error by: %s", err)
// 		return
// 	}

// 	fmt.Println(results)

// 	for results.Next() {
// 		err = results.Scan(&TemperatureMin, &TemperatureMax)
// 		if err != nil {
// 			log.Fatalf("Error by: %s", err)
// 		}
// 		fmt.Println(response)
// 	}
// }

// func Connection() {
// 	instance := NewInstance()
// 	instance.DatabaseConnection()
// 	instance.GetValues()
// }
