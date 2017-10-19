package main

import (
	"database/sql"
	"fmt"
	"os"
	// _ allows us to import the MYSQL driver without creating a local name
	// for the package.
	// This ensures the package gets into your built executable,
	// but avoids the compile error you'd normally get from
	// not calling any functions within that package.
	_ "github.com/go-sql-driver/mysql"
)

// Contact represents a contact record.
type Contact struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
}

func main() {
	// Create the data source name, which identifies the
	// user, password, server address, and default database.
	dsn := fmt.Sprintf("root:%s@tcp(192.168.99.100:3306)/demo", os.Getenv("MYSQL_ROOT_PASSWORD"))

	// Create a database object, which manages a pool of
	// network connections to the database server.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		os.Exit(1)
	}

	// Ensure that the database gets closed when we are done.
	defer db.Close()

	// For now, just ping the server to ensure we have
	// a live connection to it.
	if err := db.Ping(); err != nil {
		fmt.Printf("error pinging database: %v\n", err)
	} else {
		fmt.Printf("successfully connected!\n")
	}

	// Insert a new row into the "contacts" table.
	// Use ? markers for the values to defeat SQL
	// injection attacks.
	insq := "insert into contacts(email, first_name, last_name) values (?,?,?)"
	res, err := db.Exec(insq, "test@test.com", "Test", "Tester")
	if err != nil {
		fmt.Printf("error inserting new row: %v\n", err)
	} else {
		//get the auto-assigned ID for the new row
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("error getting new ID: %v\n", id)
		} else {
			fmt.Printf("ID for new row is %d\n", id)
		}
	}

	// Select rows from the table.
	// Go doesn't buffer the entire result set in memory, as it could be enormous.
	// db.Query returns sql.Rows object, which lets us load one row at a time into memory.
	rows, err := db.Query("select id,email,first_name,last_name from contacts")
	// Ensure the rows are closed.
	defer rows.Close()

	fmt.Printf("rows\n-----\n")

	// Create a contact struct and scan
	// the columns into the fields.
	contact := Contact{}

	// Each time you call rows.Next(),
	// the internal row buffer is overwritten with the next row's data,
	// so only one row remains in memory at a time.
	// While there are more rows...
	for rows.Next() {
		// Scan each record into struct fields.
		if err := rows.Scan(&contact.ID, &contact.Email,
			&contact.FirstName, &contact.LastName); err != nil {
			fmt.Printf("error scanning row: %v\n", err)
		}

		// Print the struct values to std out.
		fmt.Printf("%d, %s, %s, %s\n", contact.ID, contact.Email,
			contact.FirstName, contact.LastName)
	}

	// If we got an error fetching the next row, report it.
	if err := rows.Err(); err != nil {
		fmt.Printf("error getting next row: %v\n", err)
	}
}
