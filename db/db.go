// db package handles all DB Operation
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Database Driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Storage Holds Everything of DB Package
type Storage interface {
	Pingdb()
}

// Con The Connection Struct
type stroage struct {
	db *gorm.DB
}

var STRG Storage

func init() {
	connection := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DBUSERNAME"), os.Getenv("DBNAME"), os.Getenv("DBPASS"))
	db, err := gorm.Open("postgres", connection)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	STRG = &stroage{db: db}

}

// Pingdb pings the db
func (con *stroage) Pingdb() {
	err := con.db.DB().Ping()
	if err != nil {
		fmt.Println("DB Connected Successfully")
	}
}
