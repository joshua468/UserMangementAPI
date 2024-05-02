package model

import "database/sql"

// User represents the user model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Migrate creates the user table if it doesn't exist
func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			age INT
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
