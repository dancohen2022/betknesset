package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dancohen2022/betknesset/models"
	"github.com/gorilla/mux"
)

func main() {
	models.InitSynagogues()
	models.InitSynagoguesPath()

	go handler()

	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()

}

func handler() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomePage).Methods("GET")
	mux.HandleFunc("/", SynagoguePage).Methods("POST")
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

		resString = fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = fmt.Sprintf("Synagogue %s", s)

		calend := s.CalendarApi
		zman := s.ZmanimApi

		UpdateDirs(name)
		//UpdateFiles(name)

		resString = resString + GetSynagogueHttpJson(calend) + "\n\n\n\n" + GetSynagogueHttpJson(zman)
	}
	res.Write([]byte(resString))

}

func SynagogueExist(name, key string) (models.Synagogue, error) {
	syn := *models.GetSynagogues()
	b := models.Synagogue{}
	for _, s := range syn {
		if (s.Name == name) && (s.Key == key) {
			fmt.Printf("Synagogue exist: %v\n", s)
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
	if !models.DirExist(name) {
		models.CreateDir(name)
	}
}

func UpdateFiles(name string) {
	//get new JSON, parse and create new files
}
