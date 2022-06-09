package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dancohen2022/betknesset/models"
	"github.com/gorilla/mux"
)

func main() {
	models.InitSynagogues()
	models.InitSynagoguesPath()

	go handler()

	//Development Process Exit
	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()

}

func handler() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomePage).Methods("GET")
	mux.HandleFunc("/", SynagoguePage).Methods("POST")
	//mux.HandleFunc("/login", UserPage).Methods("GET")
	//mux.HandleFunc("/login", UpdateUser).Methods("POST")
	mux.HandleFunc("/admin", AdminPage).Methods("GET")
	mux.HandleFunc("/admin", UpdateSynagogues).Methods("POST")

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

		calend := s.CalendarApi
		zman := s.ZmanimApi

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
		fmt.Println("Synagogue Exist and start update files")

		calend := UpdateApiParams(s.CalendarApi)
		zman := UpdateApiParams(s.ZmanimApi)
		fmt.Printf("calend: %s\n\n\n", calend)
		fmt.Printf("zman: %s\n\n\n", zman)

		UpdateDirs(name)
		UpdateFiles(name, GetSynagogueHttpJson(calend), GetSynagogueHttpJson(zman))

		resString = fmt.Sprintf("Synagogue %s has been updated", s)
	}
	res.Write([]byte(resString))

}

func UpdateApiParams(api string) string {
	fmt.Println("UpdateApiParams")
	//Files limited period
	const PERIOD int = 14

	bStart := DateFormat(time.Now().String())
	bEnd := DateFormat(time.Now().AddDate(0, 0, PERIOD).String())
	periodStart := strings.TrimSpace(fmt.Sprintf("start=%s", bStart))
	periodEnd := fmt.Sprintf("end=%s", bEnd)
	newPeriod := fmt.Sprintf("%s&%s", periodStart, periodEnd)
	newApi := strings.Replace(api, "year=now", "", 3)
	newApi = strings.TrimSpace(newApi + newPeriod)
	fmt.Printf("newApi: %s\n", newApi)

	fmt.Printf("newApi: %s\n", newApi)
	return newApi

}

func DateFormat(dateString string) string {
	//YYYY-MM-DD
	fmt.Println("DateFormat")
	fmt.Printf("dateString: %s\n\n\n", dateString)
	d := []byte(dateString)
	d = d[:11]
	s := string(d)
	fmt.Printf("s: %s\n\n\n", s)
	return s
}

func SynagogueExist(name, key string) (models.Synagogue, error) {
	fmt.Println("SynagogueExist")
	syn := *models.GetSynagogues()
	b := models.Synagogue{}
	for _, s := range syn {
		if (s.Name == name) && (s.Key == key) {
			fmt.Printf("Synagogue exist: %v\n\n\n", s)
			b.Key = s.Key
			fmt.Printf("b.Key: %v\n\n\n", b.Key)
			b.Name = s.Name
			fmt.Printf("b.Name: %v\n\n\n", b.Name)
			b.CalendarApi = s.CalendarApi
			fmt.Printf("b.CalendarApi: %v\n\n\n", b.CalendarApi)
			b.ZmanimApi = s.ZmanimApi
			fmt.Printf("b.ZmanimApi: %v\n\n\n", b.ZmanimApi)
			return b, nil
		}
	}
	return b, errors.New("Synagogue doesn't exist")
}

func GetSynagogueHttpJson(link string) string {
	fmt.Println("GetSynagogueHttpJson")
	fmt.Printf("link - %s\n", link)

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
	if !models.DirExist(name) {
		models.CreateDir(name)
	}
}

func UpdateFiles(name, calend, zman string) {
	fmt.Println("UpdateFiles")
	/*
		1. Delete all Daily files
		2. Create new Daily files
	*/

	v := models.ZmanimJson{}
	err := json.Unmarshal([]byte(zman), &v)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("v: %v\n\n", v)
	fmt.Printf("V chatzot %v\n\n", v.Times.Chatzot)

	//get new JSON, parse and create new files
}
