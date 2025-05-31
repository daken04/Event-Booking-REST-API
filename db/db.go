package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3","api.db")
	if err != nil {
		panic("Could not connect to database") 
		// Gin does not close the server, it will just log the error with this
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Description TEXT NOT NULL,
		Location TEXT NOT NULL,
		DateTime DATETIME NOT NULL,
		UserID INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}