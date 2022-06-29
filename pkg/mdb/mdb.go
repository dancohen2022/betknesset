package mdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	logo string (filename)
	background string (filename)
	*/

	// SQL statement to create a task table, with no records in it.
	sqlStmt := `
	CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, name TEXT, key TEXT, type TEXT, 
		active INTEGER,config TEXT, zmanimApi TEXT, calendarApi TEXT, logo string, background string);
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
	  synagogue_name TEXT
	  name TEXT
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/
	fmt.Println("sqlStmt = CREATE TABLE schedules")
	sqlStmt = `
CREATE TABLE schedules (id INTEGER NOT NULL PRIMARY KEY, synagogue_name TEXT, name TEXT, date TEXT, info TEXT);
`
	// Execute the SQL statement
	_, err = db.Exec(sqlStmt)

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

////// OUTPUTS
func SynagogueFromRow(row *sql.Rows) (synagogues.Synagogue, error) {

	s := synagogues.Synagogue{}

	var id int64
	var name string
	var key string
	var typ string //manager, synagogue
	var active bool
	var config string
	var zmanimApi string
	var calendarApi string
	var logo string
	var background string

	err := row.Scan(&id, &name, &key, &typ, &active, &config, &zmanimApi, &calendarApi, &logo, &background)

	if err != nil {
		log.Println(err)
		return s, err
	}

	s.User.Id = id
	s.User.Name = name
	s.User.Key = key
	s.User.UserType = typ
	s.User.Active = active
	s.Config = config

	s.ZmanimApi = zmanimApi
	s.CalendarApi = calendarApi
	s.Logo = logo
	s.Background = background

	return s, nil
}

func ConfigItemFromRow(row *sql.Rows) (synagogues.ConfigItem, error) {

	c := synagogues.ConfigItem{}

	var id int64
	var synagogue_name string
	var name string
	var date string // (2022-03-16)
	var info string //(json) //JSON with all the schedules

	err := row.Scan(&id, &synagogue_name, &name, &date, &info)

	if err != nil {
		log.Println(err)
		return c, err
	}

	_ = json.Unmarshal([]byte(info), &c)

	return c, nil
}

/////// ADD/CREATE/INSERT
// CreateManager, CreateSynagogue, CreateConfigItem

//ADD Manager (synagogue)
func CreateManager(db *sql.DB, u synagogues.User) synagogues.User {

	user := synagogues.User{}

	/* users:
	id INTEGER NOT NULL PRIMARY KEY
	name TEXT
	key TEXT
	type TEXT //manager, synagogue
	active BOOLEAN
	config TEXT (json)
	zmanimApi TEXT (json)
	calendarApi TEXT (json)
	logo TEXT
	background TEXT
	*/
	sqlStmt := `INSERT INTO users (name , key, type, active, zmanimApi, calendarApi, logo, background)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, u.Name, u.Key, "manager", true, "", "", nil, nil)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return user
	}

	return user
}

//ADD SYNAGOGUE (synagogue)
func CreateSynagogue(db *sql.DB, s synagogues.Synagogue) synagogues.Synagogue {

	synagogue := synagogues.Synagogue{}

	/* users:
	id INTEGER NOT NULL PRIMARY KEY
	name TEXT
	key TEXT
	type TEXT //manager, synagogue
	active BOOLEAN
	config TEXT (json)
	zmanimApi TEXT (json)
	calendarApi TEXT (json)
	logo TEXT
	background TEXT
	*/
	sqlStmt := `INSERT INTO users (name , key, type, active, zmanimApi, calendarApi, logo, background)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, s.User.Name, s.User.Key, "synagogue", true, s.ZmanimApi, s.CalendarApi, "", "")
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return synagogue
	}

	return synagogue
}

//CREATE schedule rom schedule list
func CreateConfigItem(db *sql.DB, synagogue_name string, c synagogues.ConfigItem) synagogues.ConfigItem {
	confItem := synagogues.ConfigItem{}

	/*schedules
	  id INTEGER NOT NULL PRIMARY KEY
	  synagogue_name TEXT
	  name TEXT
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/

	j, _ := json.Marshal(c)

	sqlStmt := `INSERT INTO schedules (synagogue_name ,name, date, info)
	VALUES(?, ?, ?)`
	_, err := db.Exec(sqlStmt, synagogue_name, c.Name, c.Date, j)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return confItem
	}

	return confItem
}

/////// GET
// GetManager, GetSynagogue,GetConfigItems

//GET Synagogues BY User name and key - return []Synagogue but only 1 item
func GetSynagogue(db *sql.DB, name string, key string, typ string, active bool) []synagogues.Synagogue {
	synagogues := []synagogues.Synagogue{}
	rows, err := db.Query(
		`
		SELECT id, name, key, type, active, config, zmanimApi, calendarApi, logo, background
		FROM users
		WHERE name = ? AND key = ? AND type = ? AND active = ?
		`, name, key, typ, active)

	if err != nil {
		log.Println(err)
		return synagogues
	}

	defer rows.Close()

	for rows.Next() {
		s, err := SynagogueFromRow(rows)
		if err != nil {
			log.Println(err)
			return synagogues
		}
		synagogues = append(synagogues, s)
	}
	return synagogues
}

//GET ConfigItems BY synagogue name and date - return ConfigItem
func GetConfigItems(db *sql.DB, synagogueName string, date string) []synagogues.ConfigItem {
	confItems := []synagogues.ConfigItem{}

	/*schedules
	  id INTEGER NOT NULL PRIMARY KEY
	  synagogue_name TEXTname TEXT
	  name TEXT
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/

	rows, err := db.Query(
		`
		SELECT id, synagogue_name, name, key, date, info
		FROM schedules
		WHERE synagogue_name = ? AND date = ?
		`, synagogueName, date)

	if err != nil {
		log.Println(err)
		return confItems
	}

	defer rows.Close()

	for rows.Next() {
		c, err := ConfigItemFromRow(rows)
		if err != nil {
			log.Println(err)
			return confItems
		}
		confItems = append(confItems, c)
	}
	return confItems
}

//GET all users (synagogues format) - return slice of users
func GetAllSynagogues(db *sql.DB) []synagogues.Synagogue {
	synagogues := []synagogues.Synagogue{}
	rows, err := db.Query(
		`
		SELECT id, name, key, type, active, config, zmanimApi, calendarApi, logo, background
		FROM users
		`)

	if err != nil {
		log.Println(err)
		return synagogues
	}

	defer rows.Close()

	for rows.Next() {
		s, err := SynagogueFromRow(rows)
		if err != nil {
			log.Println(err)
			return synagogues
		}
		synagogues = append(synagogues, s)
	}
	return synagogues
}

/////// UPDATE

// UPDATE  user and synagogue manager return user
func UpdateSynagogue(db *sql.DB, s synagogues.Synagogue) []synagogues.Synagogue {
	synagogues := []synagogues.Synagogue{}
	sqlStmt := `
	UPDATE users SET  key=?, type=?, active=?, config=?, zmanimApi=?, calendarApi=?, logo=?, background=?
	WHERE name = ?
	`
	_, err := db.Exec(sqlStmt, s.User.Key, s.User.UserType, s.User.Active, s.Config, s.ZmanimApi, s.CalendarApi, s.Logo,
		s.Background, s.User.Name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return synagogues
	}

	return GetSynagogue(db, s.User.Name, s.User.Key, s.User.UserType, s.User.Active)
}

// UPDATE  schedule BY date - return schedule

func UpdateConfigItem(db *sql.DB, synagogueName string, c synagogues.ConfigItem) []synagogues.ConfigItem {
	configItems := []synagogues.ConfigItem{}

	/*schedules
	  id INTEGER NOT NULL PRIMARY KEY
	  synagogue_name TEXT
	  name TEXT
	  date TEXT (2022-03-16)
	  info string (json) //JSON with all the schedules
	*/
	j, _ := json.Marshal(c)

	sqlStmt := `
	UPDATE schedules SET  info=?
	WHERE synagogue_name = ? AND name = ? AND date = ?
	`
	_, err := db.Exec(sqlStmt, string(j), synagogueName, c.Name, c.Date)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return configItems
	}

	return GetConfigItems(db, synagogueName, c.Date)
}

/////// DELETE
// DELETE user BY name
func DeleteUser(db *sql.DB, name string) error {
	if name == "" {
		name = "*"
	}

	sqlStmt := `
	DELETE FROM users
	WHERE name = ? 
	`
	_, err := db.Exec(sqlStmt, name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	sqlStmt = `
	DELETE FROM schedules
	WHERE synagogue_name = ? 
	`
	_, err = db.Exec(sqlStmt, name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

//DELETE schedule BY synagogue_name, name, date
//DELETE schedule BY date
func DeleteSchedules(db *sql.DB, synagogue_name, name, date string) error {
	if synagogue_name == "" {
		synagogue_name = "*"
	}
	if name == "" {
		name = "*"
	}
	if date == "" {
		date = "*"
	}

	sqlStmt := `
	DELETE FROM schedules
	WHERE synagogue_name = ? AND name = ? AND date  = ?
	`
	_, err := db.Exec(sqlStmt, synagogue_name, name, date)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}
