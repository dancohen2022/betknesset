package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dancohen2022/betknesset/pkg/mdb"
	"github.com/dancohen2022/betknesset/pkg/synagogues"
)

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
	zmn := synagogues.ZmanimJson{}

	err := json.Unmarshal([]byte(calend), &cld)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal([]byte(zman), &zmn)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		fmt.Printf("v: %v\n\n", v)
		fmt.Printf("V chatzot %v\n\n", v.Times.Chatzot)
	*/

	//get new JSON, parse and create new files
}

func GetSynagogueHttpJson(api string) []byte {
	fmt.Println("GetSynagogueHttpJson")

	resp, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	return body
}

func GetLogoName(synName string) string {
	fmt.Println("GetLogoName")
	return "logo_" + synName
}
func GetBackgroundName(synName string) string {
	fmt.Println("GetBackgroundName")
	return "background_" + synName
}

func UpdateCalendarJSON(synName, calendAPI string) {
	fmt.Println("UpdateCalendarJSON")
	jByte := GetSynagogueHttpJson(calendAPI)
	//fmt.Printf("jByte: %s\n", string(jByte))
	calJson := synagogues.CalendarJson{}
	err := json.Unmarshal(jByte, &calJson)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("calJson: %v\n", calJson)
	cnfItemList := synagogues.ParseCalendarItemsToConfigItems(calJson.Items)

	for _, item := range cnfItemList {
		mdb.CreateConfigItem(synName, item)
	}
}

func UpdateZmanimJSON(synName, zmanAPI string) {
	fmt.Println("UpdateZmanimJSON")
	jByte := GetSynagogueHttpJson(zmanAPI)
	//fmt.Printf("jByte: %s\n", string(jByte))
	zmanJson := synagogues.ZmanimJson{}
	err := json.Unmarshal(jByte, &zmanJson)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("zmanJson: %v\n", zmanJson)
	cnfItemList := synagogues.ParseZmanimJsonToConfigItems(zmanJson)
	for _, item := range cnfItemList {
		mdb.CreateConfigItem(synName, item)
	}

}

func UpdateDefaultConfigItemsList(synName string) {
	fmt.Println("UpdateDefaultConfigItemsList")
	cnfItemList := GetDefaultConfigItemsList()
	for _, item := range cnfItemList {
		mdb.CreateConfigItem(synName, item)
	}

}
