package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dancohen2022/betknesset/pkg/functions"
	"github.com/gorilla/mux"
)

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
	s, err := functions.SynagogueExist(name, key)
	if err != nil {
		resString = resString + fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = resString + "Synagogue dowsn't exist\n"
	} else {

		resString = fmt.Sprintf("<h1>name is: %s and key is: %s</h1>\n", name, key)
		resString = fmt.Sprintf("Synagogue %v", s)

		calend := functions.UpdateApiParams(s.CalendarApi)
		zman := functions.UpdateApiParams(s.ZmanimApi)
		fmt.Println(calend)
		fmt.Println(zman)

		functions.UpdateDirs(name)

		functions.UpdateSynagogueSchedule(name)

		//resString = resString + functions.GetSynagogueHttpJson(calend) + "\n\n\n\n" + functions.GetSynagogueHttpJson(zman)
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
