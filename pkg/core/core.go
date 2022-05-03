package core

import (
	"database/sql"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/database"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/migrations"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/models"
)

// CreateMySQLClient Create a new MySQL Client, so we can make a connection to a database.
func CreateMySQLClient() database.MySQLDB {
	return database.MySQLDB{}
}

// MigrateDB Given a connection string and an array of `MigrationData` scripts,
// will attempt to run the migrations against your database. If this is the first
// time you use this tool, it will auto create a `migrations` table on the database first.
// Migrations are stored in that table, and you can go query the database directly
// to check the status of a particular migration.
func MigrateDB(db *sql.DB, data []models.DBMigrationData) {
	migrations.MigrateDBExec(db, data)
}
