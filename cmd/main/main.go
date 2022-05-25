package main

import (
	"log"
	"net/http"

	routes "github.com/dancohen2022/betknesset/pkg/calendar-routes"
	"github.com/dancohen2022/betknesset/pkg/models"
	"github.com/gorilla/mux"
)

func main() {
	// Get synagogs list
	models.InitSynagogues()
	//

	r := mux.NewRouter()
	routes.RegisterBetknessetRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
