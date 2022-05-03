package models

// DBMigrationData represents the name of a migration file and the SQL we need to run.
type DBMigrationData struct {
	Name string
	File string
	SQL  string
}

// DBMigrationEvent represents a migration event that has happened that we will store
// in the database, so we can check on the next deployment if we need to run migrations.
type DBMigrationEvent struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	File    string `json:"file"`
	Created string `json:"created"`
}
