package main

import (
	"github.com/dancohen2022/betknesset/pkg/models"
	"github.com/dancohen2022/betknesset/pkg/routes"
)

func main() {
	// Get synagogs list
	models.InitSynagogues()
	//

	models.CreateDir("test")

	link := "https://www.hebcal.com/"
	json := "hebcal?v=1&cfg=json&maj=on&min=on&mod=on&nx=on&year=now&month=x&ss=on&mf=on&c=on&geo=geoname&geonameid=3448439&M=on&s=on"

	routes.GetSynagogueHttpJson(link, json)

	/*
		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			key := "q"
			val := req.URL.Query().Get(key)
			io.WriteString(res, "Do my search:"+val)
		})
		/*

		/*
			r := mux.NewRouter()
			routes.RegisterBetknessetRoutes(r)
			http.Handle("/", r)
			log.Fatal(http.ListenAndServe("localhost:9010", r))
	*/
}
