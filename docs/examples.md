# Connect/Close A MySQL Database
```go
package main

// Create a new client instance.
client := core.MySQLDB{}

// Get a connection to the database.
db := client.Connect("ConnectionString", true)

// You can also pass around the client which will also hold a connection
client.DB.Ping()

// Close the connection
err := client.Close()
if err != nil {
	panic(err)
}
```

# Perform Migrations
```go
package main

// Create a new client instance.
client := core.MySQLDB{}

// Get a connection to the database.
db := client.Connect("ConnectionString", true)

// Build Migration Array
var migrationsData = []core.DBMigrationData{
	core.DBMigrationData{
		Name: "Initial Migration",
		File: "migration-initial.go",
		SQL:  "SELECT * FROM users",
	},
}

// Run Migrations
core.MigrateDB(db, migrationsData)

client.Close()
```
