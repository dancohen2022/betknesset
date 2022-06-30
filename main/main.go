package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/dancohen2022/betknesset/pkg/mdb"
	"github.com/dancohen2022/betknesset/pkg/synagogues"
	"github.com/mattn/go-sqlite3"
)

const PERIOD int = 14

func main() {
	///// OPEN DATABASE CONNECTION
	// Remove the todo database file if exists.
	// Comment out the below line if you don't want to remove the database.
	os.Remove(mdb.SYNAGOGUESDB)
	// Open database connection
	db, err := sql.Open("sqlite3", mdb.SYNAGOGUESDB)
	// Check if database connection was opened successfully
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
	// close database connection before exiting program.
	defer db.Close()
	////// CREATE TABLES if first time
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		mdb.CreateDBTables(db)
	}()
	wg.Wait()

	u := synagogues.User{
		Id:       0,
		Name:     "Dan",
		Key:      "12345",
		UserType: "manager",
		Active:   true,
	}

	/*
		User        User   `json:"user"`
		CalendarApi string `json:"calendar"`
		ZmanimApi   string `json:"zmanim"`
		Config      string `json:"config"`
		Logo        string `json:"logo"`
		Background  string `json:"background"`
	*/

	s := []synagogues.Synagogue{
		{User: synagogues.User{
			Id:       0,
			Name:     "shuva_raanana",
			Key:      "123456",
			UserType: "synagogue",
			Active:   true,
		},
			CalendarApi: "https://www.hebcal.com/hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&ss=on&mf=on&c=on&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel&M=on&s=on",
			ZmanimApi:   "https://www.hebcal.com/zmanim?cfg=json&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel&year=now",
			Config:      "",
			Logo:        "",
			Background:  ""},
		{User: synagogues.User{
			Id:       0,
			Name:     "bentata",
			Key:      "654321",
			UserType: "synagogue",
			Active:   true,
		},
			CalendarApi: "https://www.hebcal.com/hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&ss=on&mf=on&c=on&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel&M=on&s=on",
			ZmanimApi:   "https://www.hebcal.com/zmanim?cfg=json&geo=geopos&latitude=32.1848&longitude=34.8713&tzid=Israel&year=now",
			Config:      "",
			Logo:        "",
			Background:  "",
		},
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mdb.CreateManager(db, u)
		mdb.CreateSynagogue(db, s[0])
		mdb.CreateSynagogue(db, s[1])
	}()
	wg.Wait()
	/////
	//Init from files
	//functions.InitSynagogues()
	//functions.CreatFirstDefaultConfigValuesFile()

	/////
	/*
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			manager.LoopManager()
		}()
	*/

	/*
		wg.Add(1)
		go func(){
			defer wg.Done()
			go handler()
			log.Println("Server started, press <ENTER> to exit")
			fmt.Scanln()
		}()
	*/

	//wg.Wait()

}
