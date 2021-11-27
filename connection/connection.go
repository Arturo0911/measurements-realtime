package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
