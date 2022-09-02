package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Dpyadav@123@tcp(127.0.0.1:3306)/institution")
	if err != nil {
		return nil, err
	}

	fmt.Println("database connected")

	return db, nil
}
