package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dancohen2022/betknesset"
	"github.com/gorilla/mux"
)

func main() {
	InitSynagogues()

	go handler()

	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()

	//models.CreateDir("test")
}

func handler() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", HomePage).Methods("GET")
	mux.HandleFunc("/synagogue", SynagoguePage)

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
	<form action="/synagogue" method="POST">
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

	fmt.Printf("req Method:%v\n", req.Method)
	fmt.Printf("req req.URL.Path:%s, %s\n", req.FormValue("name"), req.FormValue("key"))
	name := req.FormValue("name")
	key := req.FormValue("key")

	resString := fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)

	link := "https://www.hebcal.com/"
	json := "hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&month=x&ss=on&mf=on&c=on&geo=geoname&geonameid=3448439&M=on&s=on"

	resString = resString + GetSynagogueHttpJson(link, json)
	res.Write([]byte(resString))

}

func SynagogueExist(name, key string) {
	syn := &betknesset.GetSynagogues()
	for _, s := range syn {

	}
}

func GetSynagogueHttpJson(link, json string) string {

	resp, err := http.Get(fmt.Sprint(link, json))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//Print the HTTP response status.

	//fmt.Println("Response status:", resp.Status)
	//Print the first 5 lines of the response body.

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Body : %s", body)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	return string(body)
}
