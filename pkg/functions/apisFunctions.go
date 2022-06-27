package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
)

const PERIOD int = 14

func UpdateApiParams(api string) string {
	fmt.Println("UpdateApiParams")
	//Files limited period

	bStart := DateFormat(time.Now().String())
	bEnd := DateFormat(time.Now().AddDate(0, 0, PERIOD).String())
	periodStart := strings.TrimSpace(fmt.Sprintf("start=%s", bStart))
	periodEnd := fmt.Sprintf("end=%s", bEnd)
	newPeriod := fmt.Sprintf("&%s&%s", periodStart, periodEnd)
	newApi := strings.Replace(api, "&year=now", "", 3)
	newApiTrimes := strings.TrimSpace(newApi + newPeriod)

	return newApiTrimes

}

func DateFormat(dateString string) string {
	//YYYY-MM-DD
	fmt.Println("DateFormat")
	d := []byte(dateString)
	d = d[:11]
	s := string(d)
	return s
}

func UpdateDirs(name string) {
	fmt.Println("UpdateDirs")
	if !DirExist(name) {
		CreateDir(name)
	}
}

func UpdateFiles(name, calend, zman, config string) {
	fmt.Println("UpdateFiles")
	/*
		1. Delete all Daily files
		2. Create new Daily files
	*/
	cld := synagogues.CalendarJson{}
	zmn := functions.ZmanimJson{}
	cnf := functions.ConfigJson{}

	err := json.Unmarshal([]byte(calend), &cld)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(zman), &zmn)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(config), &cnf)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		fmt.Printf("v: %v\n\n", v)
		fmt.Printf("V chatzot %v\n\n", v.Times.Chatzot)
	*/

	//get new JSON, parse and create new files
}

func GetSynagogueHttpJson(link string) string {
	fmt.Println("GetSynagogueHttpJson")

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

func GetSynagogueConfigJson(name string) string {
	fmt.Println("GetSynagogueConfigJson")

	resp := ReadFile(name, functions.CONFIGFILE)
	return string(*resp)

}
