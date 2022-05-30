package routes

import (
	"github.com/dancohen2022/betknesset/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBetknessetRoutes = func(router *mux.Router) {
	router.HandleFunc("/betknesset/{name, key}", controllers.GetDayTimes).Methods("GET")

}
