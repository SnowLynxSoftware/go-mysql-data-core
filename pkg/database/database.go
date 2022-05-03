package database

import (
	"database/sql"
	"fmt"
	"github.com/SnowLynxSoftware/go-mysql-data-core/configs"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLDB Defines a database connection
type MySQLDB struct {
	ConnectionString string
	DB               *sql.DB
}

func (mysqlDB MySQLDB) Connect(connectionString string, allowMultiStatements bool) (*sql.DB, error) {
	mysqlDB.ConnectionString = connectionString
	db, err := initializeDatabaseConnectionExec(connectionString, allowMultiStatements)
	if db != nil {
		mysqlDB.DB = db
		return mysqlDB.DB, nil
	} else {
		return nil, err
	}
}

func (mysqlDB MySQLDB) CloseConnection() error {
	err := mysqlDB.DB.Close()
	return err
}

func initializeDatabaseConnectionExec(connectionString string, allowMultiStatements bool) (*sql.DB, error) {
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
		return nil, err
	}

	// Ping the server to make sure we have a good connection
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[go-mysql-data-core " + configs.GetVersion() + "] Database connection is open!")
	}

	return db, nil
}
