// db/db.go
package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// +--------------------------------------------+
// | laptopID = PRIMARY KEY                     |
// | All other tables use this as a FOREIGN KEY |
// +--------------------------------------------+

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./laptop_tracker.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create Laptops table
	createLaptopsTable := `
    CREATE TABLE IF NOT EXISTS laptops (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        donor_name TEXT,
        specs TEXT,
        status TEXT,
        donation_date TEXT
    );`
	_, err = db.Exec(createLaptopsTable)
	if err != nil {
		log.Fatal("Error creating laptops table:", err)
	}

	// Create Repairs table
	createRepairsTable := `
    CREATE TABLE IF NOT EXISTS repairs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        laptop_id INTEGER,
        issue TEXT,
        technician TEXT,
        status TEXT,
				reasoningForRepair TEXT,
        FOREIGN KEY(laptop_id) REFERENCES laptops(id)
    );`
	_, err = db.Exec(createRepairsTable)
	if err != nil {
		log.Fatal("Error creating repairs table:", err)
	}

	// Create Recipients table
	createRecipientsTable := `
    CREATE TABLE IF NOT EXISTS recipients (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        contact_info TEXT,
        laptop_id INTEGER,
        received_at TEXT,
        FOREIGN KEY(laptop_id) REFERENCES laptops(id)
    );`
	_, err = db.Exec(createRecipientsTable)
	if err != nil {
		log.Fatal("Error creating recipients table:", err)
	}

	createPhonesTable := `
	CREATE TABLE IF NOT EXISTS phones (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		donor_name TEXT,
		specs TEXT,
		status TEXT,
		donation_date TEXT 
	);`
	_, err = db.Exec(createPhonesTable)
	if err != nil {
		log.Fatal("Error creating phones table", err)
	}
	return db
}
