package mdb

import (
	"database/sql"
	"log"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
	_ "github.com/mattn/go-sqlite3"
)

//sqlite cruds

const SYNAGOGUESDB = "synagogues.db"

func CreateDBTables(db *sql.DB) {

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
	  name TEXT
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/
	sqlStmt = `
CREATE TABLE schedules (id INTEGER NOT NULL PRIMARY KEY, name TEXT, date TEXT, json TEXT);
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

/////// ADD

//ADD SYNAGOGUE (synagogue)
func AddSynagogue(db *sql.DB, s synagogues.Synagogue) synagogues.Synagogue {
	synagogue := synagogues.Synagogue{}
	sqlStmt := `
CREATE TABLE schedules (id INTEGER NOT NULL PRIMARY KEY, name TEXT, date TEXT, json TEXT);
`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return synagogue
	}

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
