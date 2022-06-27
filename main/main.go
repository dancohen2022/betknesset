package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dancohen2022/betknesset/pkg/functions"
	"github.com/dancohen2022/betknesset/pkg/manager"
	"github.com/dancohen2022/betknesset/pkg/synagogues"
	"github.com/gorilla/mux"
)

const PERIOD int = 14

func main() {
	functions.InitSynagogues()

	functions.CreatFirstDefaultConfigValuesFile()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		manager.LoopManager()
	}()
	/*
		wg.Add(1)
		go func(){
			defer wg.Done()
			go handler()
			log.Println("Server started, press <ENTER> to exit")
			fmt.Scanln()
		}()
	*/

	wg.Wait()

}

func handler() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomePage).Methods("GET")
	mux.HandleFunc("/", SynagoguePage).Methods("POST")
	//mux.HandleFunc("/login", UserPage).Methods("GET")
	//mux.HandleFunc("/login", UpdateUser).Methods("POST")
	mux.HandleFunc("/admin", AdminPage).Methods("GET")
	//mux.HandleFunc("/admin", UpdateSynagogues).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("req Method:%v\n", req.Method)
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	//res.Header().Add("Content-Type", "application/json")
	body := `
	<!DOCTYPE html>
	<html>
	<body>
	<form action="/" method="POST">
	<input id="name" name="name" type="text" placeholder="name">
	<input id="key" name="key" type="text" placeholder="key">
	<input type="submit" value="Get Page">
	</form>
	</body>
	</html>`
	res.Write([]byte(body))

}

func SynagoguePage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	name := req.FormValue("name")
	key := req.FormValue("key")
	resString := ""
	s, err := SynagogueExist(name, key)
	if err != nil {
		resString = resString + fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = resString + "Synagogue dowsn't exist\n"
	} else {

		resString = fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = fmt.Sprintf("Synagogue %s", s)

		calend := UpdateApiParams(s.CalendarApi)
		zman := UpdateApiParams(s.ZmanimApi)

		UpdateDirs(name)

		UpdateFiles(name, GetSynagogueHttpJson(calend), GetSynagogueHttpJson(zman))

		resString = resString + GetSynagogueHttpJson(calend) + "\n\n\n\n" + GetSynagogueHttpJson(zman)
	}
	res.Write([]byte(resString))

}

func AdminPage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	body := `
		<!DOCTYPE html>
		<html>
		<body>
		<form action="/admin" method="POST">
		<input id="name" name="name" type="text" placeholder="name">
		<input id="key" name="key" type="text" placeholder="key">
		<input type="submit" value="Update Page">
		</form>
		</body>
		</html>`
	res.Write([]byte(body))

}

/*
func UpdateSynagogues(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	name := req.FormValue("name")
	key := req.FormValue("key")
	resString := ""
	s, err := SynagogueExist(name, key)
	if err != nil {
		resString = resString + fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = resString + "Synagogue dowsn't exist\n"
	} else {
		fmt.Printf("Synagogue Exist and start update files\n\n\n")
		cal_api := s.CalendarApi
		zmn_api := s.ZmanimApi
		calend := UpdateApiParams(cal_api)
		zman := UpdateApiParams(zmn_api)
		fmt.Printf("calend: %s\n\n\n", calend)
		fmt.Printf("zman: %s\n\n\n", zman)

		fmt.Printf("name: %s\n\n\n", name)
		UpdateDirs(name)

		UpdateFiles(name, GetSynagogueHttpJson(calend), GetSynagogueHttpJson(zman))

		resString = fmt.Sprintf("Synagogue %s has been updated", s)
	}
	res.Write([]byte(resString))

}
*/

func UpdateApiParams(api string) string {
	fmt.Println("UpdateApiParams")
	//Files limited period

	bStart := DateFormat(time.Now().String())
	bEnd := DateFormat(time.Now().AddDate(0, 0, PERIOD).String())
	periodStart := strings.TrimSpace(fmt.Sprintf("start=%s", bStart))
	periodEnd := fmt.Sprintf("end=%s", bEnd)
	newPeriod := fmt.Sprintf("&%s&%s", periodStart, periodEnd)
	newApi := strings.Replace(api, "&year=now", "", 3)
	newApiTrimes := strings.TrimSpace(newApi + newPeriod)

	return newApiTrimes

}

func DateFormat(dateString string) string {
	//YYYY-MM-DD
	fmt.Println("DateFormat")
	d := []byte(dateString)
	d = d[:11]
	s := string(d)
	return s
}

func SynagogueExist(name, key string) (synagogues.Synagogue, error) {
	fmt.Println("SynagogueExist")
	syn := *functions.GetSynagogues()
	b := synagogues.Synagogue{}
	for _, s := range syn {
		if (s.Name == name) && (s.Key == key) {
			b.Key = s.Key
			b.Name = s.Name
			b.CalendarApi = s.CalendarApi
			b.ZmanimApi = s.ZmanimApi
			return b, nil
		}
	}
	return b, errors.New("Synagogue doesn't exist")
}

func GetSynagogueHttpJson(link string) string {
	fmt.Println("GetSynagogueHttpJson")

	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	return string(body)
}

func UpdateDirs(name string) {
	fmt.Println("UpdateDirs")
	if !functions.DirExist(name) {
		functions.CreateDir(name)
	}
}

func UpdateFiles(name, calend, zman string) {
	fmt.Println("UpdateFiles")
	/*
		1. Delete all Daily files
		2. Create new Daily files
	*/

	v := synagogues.ZmanimJson{}
	err := json.Unmarshal([]byte(zman), &v)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("v: %v\n\n", v)
	fmt.Printf("V chatzot %v\n\n", v.Times.Chatzot)

	//get new JSON, parse and create new files
}
