package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dancohen2022/betknesset/pkg/db"
)

const PERIOD int = 14

func main() {
	///// OPEN DATABASE CONNECTION
	// Remove the todo database file if exists.
	// Comment out the below line if you don't want to remove the database.
	os.Remove(db.SYNAGOGUESDB)
	// Open database connection
	db, err := sql.Open("sqlite3", db.SYNAGOGUESDB)
	// Check if database connection was opened successfully
	if err != nil {
		// Print error and exit if there was problem opening connection.
		log.Fatal(err)
		return
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
