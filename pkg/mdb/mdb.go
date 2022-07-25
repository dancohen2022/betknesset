package mdb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

//sqlite cruds

const SYNAGOGUESDB = "synagogues.db"

var db *sql.DB

func GetDb() *sql.DB {
	fmt.Println("GetDb")
	return db
}
func SetDb(d *sql.DB) {
	fmt.Println("SetDb")
	db = d
}

func CreateDbTables() {
	fmt.Println("CreateDbTables")

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
	logo TEXT (filename)
	background TEXT (filename)
	*/

	// Execute the SQL statement
	// SQL statement to create a task table, with no records in it.
	sqlStmt := `
			CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, name TEXT, key TEXT, type TEXT,
				active BOOLEAN,config TEXT, zmanimApi TEXT, calendarApi TEXT, logo TEXT, background TEXT);
			`

	//`DELETE FROM users;
	//`
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
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active       bool
	*/
	//fmt.Println("sqlStmt = CREATE TABLE schedules")
	sqlStmt = `
		CREATE TABLE schedules (id INTEGER NOT NULL PRIMARY KEY, synagogue_name TEXT, 
			name TEXT, hname TEXT,category TEXT, subcat TEXT, date TEXT, time TEXT,
			info TEXT, active BOOLEAN);
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
	//fmt.Println("sqlStmt = CREATE TABLE schedules SUCCEEDED")

	//`
	//DELETE FROM schedules;
	//`
}

////// OUTPUTS
func SynagogueFromRow(row *sql.Rows) (synagogues.Synagogue, error) {
	fmt.Println("SynagogueFromRow")

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
	fmt.Println("ConfigItemFromRow")

	c := synagogues.ConfigItem{}

	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active      bool
	*/
	var id int64
	var synagogue_name string
	var name string
	var hname string
	var category string
	var subcat string
	var date string // (2022-03-16)
	var time string
	var info string //(json) //JSON with all the schedules
	var active bool

	err := row.Scan(&id, &synagogue_name, &name, &hname, &category, &subcat, &date, &time, &info, &active)

	if err != nil {
		log.Println(err)
		return c, err
	}

	c.Name = name
	c.Hname = hname
	c.Category = category
	c.Subcat = subcat
	c.Date = date
	c.Time = time
	c.Info = info
	c.Active = active

	return c, nil
}

/////// ADD/CREATE/INSERT
// CreateManager, CreateSynagogue, CreateConfigItem

//ADD Manager (synagogue)
func CreateManager(u synagogues.User) synagogues.User {
	fmt.Println("CreateManager")

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
	sqlStmt := `INSERT INTO users (name , key, type, config, active, zmanimApi, calendarApi, logo, background)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?,?)`
	_, err := db.Exec(sqlStmt, u.Name, u.Key, "manager", "", true, "", "", "", "")
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return user
	}

	return user
}

//ADD SYNAGOGUE (synagogue)
func CreateSynagogue(s synagogues.Synagogue) synagogues.Synagogue {
	fmt.Println("CreateSynagogue")

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
	sqlStmt := `INSERT INTO users (name , key, type, active, config, zmanimApi, calendarApi, logo, background)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, s.User.Name, s.User.Key, "synagogue", true, "", s.ZmanimApi, s.CalendarApi, "", "")
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return synagogue
	}

	return synagogue
}

//CREATE schedule rom schedule list
func CreateConfigItem(synagogue_name string, c synagogues.ConfigItem) synagogues.ConfigItem {
	fmt.Println("CreateConfigItem")
	//fmt.Println(c)
	confItem := synagogues.ConfigItem{}

	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active      bool
	*/

	sqlStmt := `INSERT INTO schedules (synagogue_name , name , hname ,category ,subcat , date , time , info ,active) 
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStmt, synagogue_name, c.Name, c.Hname, c.Category, c.Subcat, c.Date, c.Time, c.Info, c.Active)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return confItem
	}

	return confItem
}

/////// GET
// GetManager, GetSynagogue,GetConfigItems

//GET Synagogues BY User name and key - return []Synagogue but only 1 item
func GetSynagogue(name string, key string, typ string, active bool) []synagogues.Synagogue {
	fmt.Println("GetSynagogue")
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

//GET all users (synagogues format) - return slice of users
func GetAllSynagogues() []synagogues.Synagogue {
	fmt.Println("GetAllSynagogues")
	synagogues := []synagogues.Synagogue{}
	rows, err := db.Query(
		`
		SELECT id, name, key, type, active, config, zmanimApi, calendarApi, logo, background
		FROM users;
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

//GET ConfigItems BY synagogue name and date - return ConfigItem
func GetConfigItems(synagogueName string, date string) []synagogues.ConfigItem {
	fmt.Println("GetConfigItems")
	confItems := []synagogues.ConfigItem{}

	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active      bool
	*/

	rows, err := db.Query(
		`
		SELECT id, synagogue_name , name , hname ,category ,subcat , date , time, 
		info ,active
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

//GET ConfigItems BY synagogue name - return ConfigItem
func GetAllConfigItems(synagogueName string) []synagogues.ConfigItem {
	fmt.Println("GetAllConfigItems")
	confItems := []synagogues.ConfigItem{}

	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active      bool
	*/

	rows, err := db.Query(
		`
		SELECT id, synagogue_name , name , hname ,category ,
		subcat , date , time,  info ,active
		FROM schedules
		WHERE synagogue_name = ?
		`, synagogueName)

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

/////// UPDATE

// UPDATE  user and synagogue manager return user
func UpdateSynagogue(s synagogues.Synagogue) []synagogues.Synagogue {
	fmt.Println("UpdateSynagogue")
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

	return GetSynagogue(s.User.Name, s.User.Key, s.User.UserType, s.User.Active)
}

// UPDATE  schedule BY date - return schedule

func UpdateConfigItem(synagogueName string, c synagogues.ConfigItem) []synagogues.ConfigItem {
	fmt.Println("UpdateConfigItem")
	configItems := []synagogues.ConfigItem{}

	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active      bool
	*/

	sqlStmt := `
	UPDATE schedules SET  info=?
	WHERE synagogue_name = ? AND name = ? AND date = ?
	`
	_, err := db.Exec(sqlStmt, c.Info, synagogueName, c.Name, c.Date)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return configItems
	}

	return GetConfigItems(synagogueName, c.Date)
}

/////// DELETE
// DELETE user BY name
func DeleteUser(name string) error {
	fmt.Println("DeleteUser")
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

func DeleteSchedules(synagogue_name, name, date string) error {
	fmt.Println("DeleteSchedules")

	//DELETE schedule BY synagogue_name, name, date
	//DELETE schedule BY date
	/*schedules
	id INTEGER NOT NULL PRIMARY KEY
	synagogue_name TEXT
	Name     TEXT
	Hname    TEXT
	Category TEXT
	Subcat   TEXT
	Date     TEXT
	Time     TEXT
	Info     TEXT
	Active       bool
	*/

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

func DeleteAllSchedulesByName(synagogue_name, name string) error {
	fmt.Println("DeleteAllSchedulesByName")

	sqlStmt := `
	DELETE FROM schedules
	WHERE synagogue_name = ? AND name = ?
	`
	_, err := db.Exec(sqlStmt, synagogue_name, name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

func DeleteAllSchedulesByDate(synagogue_name, date string) error {
	fmt.Println("DeleteAllSchedulesByDate")

	sqlStmt := `
	DELETE FROM schedules
	WHERE synagogue_name = ? AND date  = ?
	`
	_, err := db.Exec(sqlStmt, synagogue_name, date)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}

func DeleteAllSynagogueSchedules(synagogue_name string) error {
	fmt.Println("DeleteAllSynagogueSchedules")

	sqlStmt := `
	DELETE FROM schedules
	WHERE synagogue_name = ?
	`
	_, err := db.Exec(sqlStmt, synagogue_name)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return nil
}
