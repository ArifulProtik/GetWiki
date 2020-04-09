package db

import "fmt"

// Pingdb pings the db
func (con *stroage) Pingdb() {
	err := con.db.DB().Ping()
	if err != nil {
		fmt.Println("DB Connected Successfully")
	}
}
