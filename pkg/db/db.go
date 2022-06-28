package db

import (
	"database/sql"
	"log"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
	_ "github.com/mattn/go-sqlite3"
)

//sqlite crud

const SYNAGOGUESDB = "synagogues.db"

func GetDb() string {
	return SYNAGOGUESDB
}

func CreateDB(db *sql.DB) {

	// both manager and synagogues are saved in the same users table

	//Create Tables
	/* users:
	id INTEGER NOT NULL PRIMARY KEY
	name TEXT
	key TEXT
	type TEXT //manager, synagogue
	active INTEGER
	config TEXT (json)
	zmanimApi TEXT (json)
	calendarApi TEXT (json)
	logo BLOB
	background BLOB
	*/
	// SQL statement to create a task table, with no records in it.
	sqlStmt := `
	CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, name TEXT, key TEXT, type TEXT, active INTEGER,config TEXT, zmanimApi TEXT, calendarApi TEXT, logo BLOB, background BLOB);
	`

	//`DELETE FROM users;
	//`

	// Execute the SQL statement
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	/*schedules
	  id INTEGER NOT NULL PRIMARY KEY
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/
	sqlStmt = `
CREATE TABLE schedules (id INTEGER NOT NULL PRIMARY KEY, date TEXT, json TEXT);
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	//`
	//DELETE FROM schedules;
	//`
}

/////// CREATE

//CREATE user (synagogue)
func CreateUser(s synagogues.Synagogue) synagogues.Synagogue {
	var synagogue synagogues.Synagogue

	db, err := sql.Open("sqlite3", SYNAGOGUESDB)

	// Check if database connection was opened successfully
	if err != nil {
		// Print error and exit if there was problem opening connection.
		log.Fatal(err)
	}
	// close database connection before exiting program.
	defer db.Close()

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Prepare prepared statement that can be reused.
	stmt, err := tx.Prepare("INSERT INTO task(id, task, owner, checked) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	// close statement before exiting program.
	defer stmt.Close()

	return synagogue
}

//CREATE schedule

/////// GET
//GET user BY name - return user

//GET all users - return slice of users

//GET all schedules BY date - return slice of schedules

//GET all schedule - return slice of schedules

/////// UPDATE

// UPDATE  user BY name - return user

// UPDATE  schedule BY date - return schedule

/////// DELETE

// DELETE user BY id (active as false)

//DELETE schedule BY date
