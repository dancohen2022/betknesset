package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Synagogue struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	CalendarApi string `json:"calendar"`
	ZmanimApi   string `json:"zmanim"`
}

type CalendarJson struct {
	Title    string           `json:"title"`
	Date     string           `json:"date"`
	Location CalendarLocation `json:"location"`
	Range    CalendarRange    `json:"range"`
	Items    []CalendarItems  `json:"items"`
}

type CalendarLocation struct {
	Title     string  `json:"title"`
	City      string  `json:"city"`
	Tzid      string  `json:"tzid"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CalendarRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type CalendarItems struct {
	Title     string          `json:"title"`
	Date      string          `json:"date"`
	Hdate     string          `json:"hdate"`
	Category  string          `json:"category"` //hebdate, omer, roshchodesh, candles, holiday, parashat, havdalah, mevarchim, zmanim
	Hebrew    string          `json:"hebrew"`
	Memo      string          `json:"memo"`
	Leyning   CalendarLeyning `json:"leyning"`
	Link      string          `json:"link"`
	TitleOrig string          `json:"title_orig"`
	Subcat    string          `json:"subcat"` //fast
	Yomtov    bool            `json:"yomtov"`
	Omer      CalendarOmer    `json:"omer"`
}

type CalendarOmer struct {
	Count  CalendarOmerCount `json:"count"`
	Sefira CalendarOmeSefira `json:"sefira"`
}

type CalendarOmerCount struct {
	He string `json:"he"`
	En string `json:"en"`
}

type CalendarOmeSefira struct {
	He       string `json:"he"`
	Translit string `json:"translit"`
	En       string `json:"en"`
}

type CalendarLeyning struct {
	L1       string `json:"1"`
	L2       string `json:"2"`
	L3       string `json:"3"`
	L4       string `json:"4"`
	L5       string `json:"5"`
	L6       string `json:"6"`
	L7       string `json:"7"`
	Torah    string `json:"torah"`
	Haftarah string `json:"haftarah"`
	Maftir   string `json:"maftir"`
}

type ZmanimJson struct {
	Date     ZmanimDate     `json:"date"`
	Location ZmanimLocation `json:"location"`
	Times    ZmanimTimes    `json:"times"`
}

type ZmanimDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type ZmanimLocation struct {
	Name      string  `json:"name"`
	Il        bool    `json:"il"`
	Tzid      string  `json:"tzid"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ZmanimTimes struct {
	ChatzotNight      interface{} `json:"chatzotNight"`
	AlotHaShachar     interface{} `json:"alotHaShachar"`
	Misheyakir        interface{} `json:"misheyakir"`
	MisheyakirMachmir interface{} `json:"misheyakirMachmir"`
	Dawn              interface{} `json:"dawn"`
	Sunrise           interface{} `json:"sunrise"`
	SofZmanShma       interface{} `json:"sofZmanShma"`
	SofZmanShmaMGA    interface{} `json:"sofZmanShmaMGA"`
	SofZmanTfilla     interface{} `json:"sofZmanTfilla"`
	SofZmanTfillaMGA  interface{} `json:"sofZmanTfillaMGA"`
	Chatzot           interface{} `json:"chatzot"`
	MinchaGedola      interface{} `json:"minchaGedola"`
	MinchaKetana      interface{} `json:"minchaKetana"`
	PlagHaMincha      interface{} `json:"plagHaMincha"`
	Sunset            interface{} `json:"sunset"`
	Dusk              interface{} `json:"dusk"`
	Tzeit7083deg      interface{} `json:"tzeit7083deg"`
	Tzeit85deg        interface{} `json:"tzeit85deg"`
	Tzeit42min        interface{} `json:"tzeit42min"`
	Tzeit50min        interface{} `json:"tzeit50min"`
	Tzeit72min        interface{} `json:"tzeit72min"`
}
type ZmanimTimeItem struct {
	ShortDate string
	LongDate  string
}

var synagogues []Synagogue

func InitSynagogues() {
	filename := "./files/synagogues/synagogues.txt"
	bFile2, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bFile2, &synagogues)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Synagogues: %v\n\n", synagogues)
}

func GetSynagogues() *[]Synagogue {
	fmt.Printf("synagogues: %v", synagogues)
	return &synagogues
}
