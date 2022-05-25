package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dancohen2022/betknesset/pkg/models"
)

var Synagogue models.Synagogue

func GetDayTimes(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Printf("body: %v\n", body)
	res := models.GetTimesJsonByNamePassword(string(body), string(body))
	w.WriteHeader(http.StatusOK)
	w.Write(json.RawMessage(*res))
}
