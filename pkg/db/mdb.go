package mdb

import (
	"database/sql"
	"log"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
	"github.com/mattn/go-sqlite3"
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
		if sqlError, ok := err.(sqlite3.Error); ok {
			// code 1 == "table already exists"
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else {
			log.Fatal(err)
		}
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

	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			// code 1 == "table already exists"
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else {
			log.Fatal(err)
		}
	}

	//`
	//DELETE FROM schedules;
	//`
}

/////// ADD/CREATE/INSERT

//ADD SYNAGOGUE (synagogue)
func CreateSynagogue(db *sql.DB, s synagogues.Synagogue) synagogues.Synagogue {

	synagogue := synagogues.Synagogue{}

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
	sqlStmt := `INSERT INTO users (name , key, type, active, config, zmanimApi, calendarApi, logo, background)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, s.User.Name, s.User.Key, "synagogue", 1, "", s.ZmanimApi, s.CalendarApi, nil, nil)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return synagogue
	}

	return synagogue
}

//CREATE schedule

/////// GET
//GET synagogue BY User - return synagogue
func GetSynagogue(db *sql.DB, user synagogues.User) synagogues.Synagogue {
	synagogue := synagogues.Synagogue{}

	return synagogue
}

//GET all users - return slice of users

//GET all schedules BY date - return slice of schedules

//GET all schedule - return slice of schedules

/////// UPDATE

// UPDATE  user BY name - return user

// UPDATE  schedule BY date - return schedule

/////// DELETE

// DELETE user BY id (active as false)

//DELETE schedule BY date
