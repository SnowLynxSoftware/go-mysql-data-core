package migrations

import (
	"database/sql"
	"fmt"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/models"
	"strconv"
)

func MigrateDBExec(db *sql.DB, dbName string, data []models.DBMigrationData) {
	checkIfMigrationTableExists(db, dbName)

	// Get all current migration events
	results, err := db.Query("SELECT * from migrations")
	if err != nil {
		panic(err)
	}

	var events []models.DBMigrationEvent

	for results.Next() {
		var event models.DBMigrationEvent

		err = results.Scan(&event.ID, &event.Name, &event.File, &event.Created)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}

	migrationsRan := 0

	// Loop through each migration data and see if we need to run that migration in this environment
	for i := 0; i < len(data); i++ {

		exists := false

		for _, v := range events {
			if data[i].Name == v.Name {
				exists = true
				break
			}
		}

		if !exists {
			_, err := db.Exec(data[i].SQL)
			if err != nil {
				fmt.Println("[MIGRATION ERROR] - " + data[i].Name)
				fmt.Printf(err.Error())
			} else {
				// Then write the new migration to the database, so we don't run it next time.
				var insertSQL = fmt.Sprintf("INSERT INTO migrations(name, file) VALUES('%s', '%s')", data[i].Name, data[i].File)
				_, err = db.Exec(insertSQL)
				if err != nil {
					return
				}
				migrationsRan++
			}
		}
	}

	defer func(db *sql.DB) {
		err := db.Close()
		fmt.Println("Database Connection Closed!")
		if err != nil {
			panic(err)
		}
	}(db)

	if migrationsRan == 0 {
		fmt.Println("No migrations necessary!")
	} else {
		fmt.Println("Migrations Ran Successfully: [" + strconv.Itoa(migrationsRan) + "]")
	}
}

func checkIfMigrationTableExists(db *sql.DB, dbName string) {
	results, err := db.Query(fmt.Sprintf("SELECT 1 FROM information_schema.TABLES WHERE TABLE_NAME = 'migrations' AND TABLE_SCHEMA = '%s';", dbName))
	tableNames := 0
	if err != nil {
		panic(err)
	}
	for results.Next() {
		tableNames++
	}
	if tableNames == 0 {
		// If there are no results returned--then we need to create the migrations table.
		var createMigrationsTableSql = `
			CREATE TABLE migrations(
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(64) NOT NULL,
				file VARCHAR(64) NOT NULL,
				created DATETIME DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY ( id )
			);`
		_, err := db.Exec(createMigrationsTableSql)
		if err != nil {
			fmt.Println("[MIGRATION TABLE CREATION ERROR] - An error occurred when attempting to initialize the migrations table!")
			fmt.Printf(err.Error())
			panic(err)
		}
	}
}
