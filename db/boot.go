// db package handles all DB Operation
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// Database Driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Con The Connection Struct
type stroage struct {
	db *gorm.DB
}

var STRG Storage

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connection := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DBUSERNAME"), os.Getenv("DBNAME"), os.Getenv("DBPASS"))
	db, err := gorm.Open("postgres", connection)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	STRG = &stroage{db: db}

}
