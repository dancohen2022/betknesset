package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dancohen2022/betknesset/pkg/mdb"
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
	//////

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
