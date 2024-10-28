package database

import (
	"database/sql"
	"devbookAPI/src/config"

	_ "github.com/go-sql-driver/mysql" // Used from sql.Open() to recognize mysql
)

// Represent the database connection
var Conn *sql.DB

// Connect(): Connect your app to the database
func Connect() (err error) {
	Conn, err = sql.Open("mysql", config.DB_URI)
	if err != nil {
		return err
	}

	if err = Conn.Ping(); err != nil {
		return err
	}

	return nil
}
