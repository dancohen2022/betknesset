package main

import (
	"github.com/dancohen2022/betknesset/pkg/models"
)

func main() {
	// Get synagogs list
	models.InitSynagogues()
	//

	models.CreateDir("test")
	/*
		r := mux.NewRouter()
		routes.RegisterBetknessetRoutes(r)
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe("localhost:9010", r))
	*/
}
