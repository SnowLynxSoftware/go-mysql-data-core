package core

import (
	"database/sql"
	"github.com/SnowLynxSoftware/go-mysql-data-core/internal/database"
	"github.com/SnowLynxSoftware/go-mysql-data-core/internal/migrations"
)

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

// InitializeDatabaseConnection Given a connection string, will attempt to return a connection
// to a MySQL Database. You can optionally pass in `allowMultiStatements` so you can run SQL
// queries in batches.
func InitializeDatabaseConnection(connectionString string, allowMultiStatements bool) *sql.DB {
	return database.InitializeDatabaseConnectionExec(connectionString, allowMultiStatements)
}

// MigrateDB Given a connection string and an array of `MigrationData` scripts,
// will attempt to run the migrations against your database. If this is the first
// time you use this tool, it will auto create a `migrations` table on the database first.
// Migrations are stored in that table, and you can go query the database directly
// to check the status of a particular migration.
func MigrateDB(connectionString string, data []DBMigrationData) {
	migrations.MigrateDBExec(connectionString, data)
}
