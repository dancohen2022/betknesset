package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	mdb.SetDb(db)
	defer db.Close()
	testMdbFunctions()
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

func testMdbFunctions() {

	/*User
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Key      string `json:"key"`
	UserType string `json:"type"` //manager, synagogue
	Active   bool   `json:"active"`

	*/
	u := synagogues.User{
		Id:       0,
		Name:     "Dan",
		Key:      "12345",
		UserType: "manager",
		Active:   true,
	}

	/* Synagogue
	User        User   `json:"user"`
	CalendarApi string `json:"calendar"`
	ZmanimApi   string `json:"zmanim"`
	Config      string `json:"config"`
	Logo        string `json:"logo"`
	Background  string `json:"background"`
	*/

	sList := []synagogues.Synagogue{
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

	/* ConfigItem
	Name     string `json:"name"`
	Hname    string `json:"hname"`
	Category string `json:"category"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Info     string `json:"info"`
	On       bool   `json:"on"`
	*/
	clist := []synagogues.ConfigItem{
		{
			Name:     "minha",
			Hname:    "מנחה",
			Category: "תפילה",
			Date:     "",
			Time:     "18:00",
			Info:     "",
			On:       true,
		},
		{
			Name:     "minha hag",
			Hname:    "מנחה של חג",
			Category: "תפילה",
			Date:     "2022-07-22",
			Time:     "17:00",
			Info:     "",
			On:       true,
		},
	}

	fmt.Print("Step 1 - CreateDBTables \n\n")
	mdb.CreateDbTables()

	fmt.Print("Step 2 - CreateManager \n\n")
	mdb.CreateManager(u)

	fmt.Print("Step 3 - CreateSynagogue \n\n")
	for _, item := range sList {
		mdb.CreateSynagogue(item)
	}

	fmt.Print("Step 4 - CreateConfigItem \n\n")
	for _, item := range clist {
		mdb.CreateConfigItem("shuva_raanana", item)
	}

	fmt.Print("Step 5 - GetAllSynagogues \n\n")
	newSList := mdb.GetAllSynagogues()
	for _, item := range newSList {
		fmt.Println(item)
	}

	fmt.Print("Step 6 - GetSynagogue \n\n")
	fmt.Println(mdb.GetSynagogue("shuva_raanana", "123456", "synagogue", true))

	fmt.Print("Step 7 - GetConfigItems \n\n")
	fmt.Println(mdb.GetConfigItems("shuva_raanana", "2022-07-22"))
	fmt.Println(mdb.GetConfigItems("shuva_raanana", ""))
	fmt.Println(mdb.GetConfigItems("bentata", "2022-07-22"))
	fmt.Println(mdb.GetConfigItems("bentata", ""))
	fmt.Println(mdb.GetAllConfigItems("shuva_raanana"))
	fmt.Println(mdb.GetAllConfigItems("bentata"))

	sList[0].Logo = "logo"
	fmt.Print("Step 9 - UpdateSynagogue \n\n")
	fmt.Println(mdb.UpdateSynagogue(sList[0]))
	mdb.DeleteUser("bentata")
	fmt.Print("Step 10 - GetAllSynagogues \n\n")
	fmt.Println(mdb.GetAllSynagogues())
	clist[1].Time = "19:00"
	fmt.Print("Step 11 - UpdateConfigItem \n\n")
	mdb.UpdateConfigItem("shuva_raanana", clist[1])
	fmt.Print("Step 12 - GetAllConfigItems \n\n")
	fmt.Println(mdb.GetAllConfigItems("shuva_raanana"))
	//fmt.Print("Step 13 - DeleteSchedules \n\n")
	//mdb.DeleteSchedules( "shuva_raanana", "minha", "")
	fmt.Print("Step 14 - GetAllConfigItems \n\n")
	fmt.Println(mdb.GetAllConfigItems("shuva_raanana"))
	fmt.Print("Step 15 - DeleteAllSchedulesByDate \n\n")
	mdb.DeleteAllSchedulesByDate("shuva_raanana", "2022-07-22")
	fmt.Print("Step 16 - GetAllConfigItems \n\n")
	fmt.Println(mdb.GetAllConfigItems("shuva_raanana"))
	fmt.Print("Step 17 - DeleteAllSchedulesByName \n\n")
	mdb.DeleteAllSchedulesByName("shuva_raanana", "minha")
	fmt.Print("Step 18 - GetAllConfigItems \n\n")
	fmt.Println(mdb.GetAllConfigItems("shuva_raanana"))

}
