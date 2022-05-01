package database

import (
	"database/sql"
	"fmt"
	"github.com/SnowLynxSoftware/go-mysql-data-core/configs"
	_ "github.com/go-sql-driver/mysql"
)

func InitializeDatabaseConnectionExec(connectionString string, allowMultiStatements bool) *sql.DB {
	// When running migrations, we should allow multi statements.
	if allowMultiStatements {
		connectionString = connectionString + "?multiStatements=true"
	}

	// Open up our database connection.
	db, err := sql.Open("mysql", connectionString)

	// If there is an error opening the connection, we will panic.
	// Database connections are important and I want to panic right away.
	if err != nil {
		fmt.Println("[go-mysql-data-core " + configs.GetVersion() + "] Error occurred when attempting a database connection!")
		panic(err.Error())
	}

	// Ping the server to make sure we have a good connection
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[go-mysql-data-core " + configs.GetVersion() + "] Database connection is open!")
	}

	return db
}
