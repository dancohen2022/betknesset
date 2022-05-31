package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dancohen2022/betknesset/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBetknessetRoutes = func(router *mux.Router) {
	router.HandleFunc("/betknesset/{name, key}", controllers.GetDayTimes).Methods("GET")

}

func GetSynagogueHttpJson(link, json string) string {
	//resp, err := http.Get("http://gobyexample.com")
	resp, err := http.Get(fmt.Sprint(link, json))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//Print the HTTP response status.

	fmt.Println("Response status:", resp.Status)
	//Print the first 5 lines of the response body.

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	return string(body)
}
