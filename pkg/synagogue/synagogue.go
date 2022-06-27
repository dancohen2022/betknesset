package synagogue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Registered Synagogues Info
type Synagogue struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	CalendarApi string `json:"calendar"`
	ZmanimApi   string `json:"zmanim"`
}

// Final Daily Schedules
type DailyScheduleJson struct {
	Title      string      `json:"title"`
	Date       string      `json:"date"`
	Hdate      string      `json:"hdate"`
	DailyItems []DailyItem `json:"dailyitems"`
}

type DailyItem struct {
	Name     string `json:"title"`
	Date     string `json:"date"`
	Hebrew   string `json:"hebrew"`
	Category string `json:"category"` //tfila, shiour
	Subcat   string `json:"subcat"`   //fast
	Time     string `json:"time"`
	Memo     string `json:"memo"`
}

// Parse Configuration File
type ConfigJson struct {
	Info    ConfigInfo   `json:"info"`
	Default []ConfigItem `json:"default"`
	Items   []ConfigItem `json:"items"`
}

type ConfigInfo struct {
	Name       string `json:"name"`
	Hname      string `json:"hname"`
	Logo       string `json:"logo"`
	Background string `json:"background"`
	Message    string `json:"message"`
}

type ConfigItem struct {
	Name     string `json:"name"`
	Hname    string `json:"hname"`
	Category string `json:"category"`
	Day      string `json:"day"`
	Time     string `json:"time"`
	Info     string `json:"info"`
	On       bool   `json:"on"`
}

// Parse Calendar API results
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

// Parse Zmanim API results
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
	// Each map has a date as key and a time  as value
	ChatzotNight      map[string]string `json:"chatzotNight"`
	AlotHaShachar     map[string]string `json:"alotHaShachar"`
	Misheyakir        map[string]string `json:"misheyakir"`
	MisheyakirMachmir map[string]string `json:"misheyakirMachmir"`
	Dawn              map[string]string `json:"dawn"`
	Sunrise           map[string]string `json:"sunrise"`
	SofZmanShma       map[string]string `json:"sofZmanShma"`
	SofZmanShmaMGA    map[string]string `json:"sofZmanShmaMGA"`
	SofZmanTfilla     map[string]string `json:"sofZmanTfilla"`
	SofZmanTfillaMGA  map[string]string `json:"sofZmanTfillaMGA"`
	Chatzot           map[string]string `json:"chatzot"`
	MinchaGedola      map[string]string `json:"minchaGedola"`
	MinchaKetana      map[string]string `json:"minchaKetana"`
	PlagHaMincha      map[string]string `json:"plagHaMincha"`
	Sunset            map[string]string `json:"sunset"`
	Dusk              map[string]string `json:"dusk"`
	Tzeit7083deg      map[string]string `json:"tzeit7083deg"`
	Tzeit85deg        map[string]string `json:"tzeit85deg"`
	Tzeit42min        map[string]string `json:"tzeit42min"`
	Tzeit50min        map[string]string `json:"tzeit50min"`
	Tzeit72min        map[string]string `json:"tzeit72min"`
}

type ZmanimTimeItem struct {
	ShortDate string
	LongDate  string
}

// End of Structures

const SYNAGOGUESPATH string = "./files/synagogues/"
const SYNAGOGUESFILE string = "./files/synagogues/synagogues.txt"
const CONFIGPATH string = "/configuration/"
const CONFIGFILE string = "config.txt"
const DEFAULTCONFIGFILE string = "./files/synagogues/defaulFilesConfig.txt"

// Synagogues List
var synagogues []Synagogue

func InitSynagogues() {

	bFile2, err := ioutil.ReadFile(SYNAGOGUESFILE)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bFile2, &synagogues)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetSynagogues() *[]Synagogue {
	fmt.Printf("synagogues: %v\n\n\n", synagogues)
	return &synagogues
}

func ResetSynagogueSchedule(name string) bool {
	for _, s := range synagogues {
		if s.Name == name {
			//Create Dir If doesnt exist yet
			UpdateDirs(name)

			//Delete all files
			if DeleteAllFiles(name) != nil {
				fmt.Println("Couldn't delete files")
				return false
			}
			//Update API period and other constants
			calend := UpdateApiParams(s.CalendarApi)
			zman := UpdateApiParams(s.ZmanimApi)

			UpdateFiles(name, GetSynagogueHttpJson(calend), GetSynagogueHttpJson(zman), GetSynagogueConfigJson(name))

			return true
		}
	}
	return false
}

func GetDaySynagogueScheduleJSON(name string) {

}

func SynagogueExist(name, key string) (Synagogue, error) {
	// Used for User Login  with Name and Key
	fmt.Println("SynagogueExist")
	syn := synagogues
	b := Synagogue{}
	for _, s := range syn {
		if (s.Name == name) && (s.Key == key) {
			b.Key = s.Key
			b.Name = s.Name
			b.CalendarApi = s.CalendarApi
			b.ZmanimApi = s.ZmanimApi
			return b, nil
		}
	}
	return b, errors.New("Synagogue doesn't exist")
}

func InitSSynagogueConfig(synagogue Synagogue) *ConfigJson {

	var config ConfigJson
	/*
		c := ReadFile(SYNAGOGUESPATH, synagogue.Name, CONFIGPATH, CONFIGPATH)
		if c == []byte{} then{
			func SetConfigDefValues(&config)

		}
		err = json.Unmarshal(bFile2, &synagogues)
		if err != nil {
			log.Fatalln(err)
		}
	*/
	return &config
}

func SetConfigDefValues(ConfigJson) *[]ConfigItem {
	s, err := ioutil.ReadFile(DEFAULTCONFIGFILE)
	c := []ConfigItem{}
	if err != nil {
		return &c
	}

	err = json.Unmarshal([]byte(s), &c)
	if err != nil {
		log.Fatalln(err)
	}
	return &c
}

func CreatFirstDefaultConfigValuesFile() {
	c := ConfigJson{}
	c.Info.Name = "Synagogue Name"
	c.Info.Hname = "שם בית כנסת"
	c.Info.Logo = "logo"
	c.Info.Background = "background"
	c.Info.Message = "about the synagogue"
	c.Default = append(c.Default, ConfigItem{Name: "WeekChacharit1", Hname: "תפילת שחרית 1", Category: "tfilot", Day: "", Time: "07:00", Info: "תפילת שחרית", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "WeekChacharit2", Hname: "תפילת שחרית 2", Category: "tfilot", Day: "", Time: "06:00", Info: "תפילת שחרית", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "WeekMincha1", Hname: "תפילת מנחה 1", Category: "tfilot", Day: "", Time: "19:00", Info: "תפילת מנחה", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "WeekMincha2", Hname: "תפילת מנחה 2", Category: "tfilot", Day: "", Time: "15:00", Info: "תפילת מנחה", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "WeekArvit1", Hname: "תפילת ערבית 1", Category: "tfilot", Day: "", Time: "20:00", Info: "תפילת ערבית", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "WeekArvit2", Hname: "תפילת ערבית 2", Category: "tfilot", Day: "", Time: "21:00", Info: "תפילת ערבית", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "ErevShabbatMincha", Hname: "תפילת מנחה ערב שבת", Category: "tfilot", Day: "", Time: "18:00", Info: "תפילת מנחה ערב שבת", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "ShabbatArvit1", Hname: "תפילת ערבית של שבת", Category: "tfilot", Day: "", Time: "18:00", Info: "תפילת ערבית של שבת", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "ShabbatChacharit", Hname: "תפילת שחרית של שבת", Category: "tfilot", Day: "", Time: "08:00", Info: "תפילת שחרית של שבת", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "ShabbatMinha1", Hname: "תפילת מנחה של שבת", Category: "tfilot", Day: "", Time: "17:00", Info: "תפילת מנחה של שבת", On: true})
	c.Default = append(c.Default, ConfigItem{Name: "ShabbatMinha2", Hname: "תפילת מנחה של שבת", Category: "tfilot", Day: "", Time: "14:00", Info: "תפילת מנחה של שבת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialChacharit1", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Day: "", Time: "08:00", Info: "תפילת שחרית מיוחדת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialChacharit2", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Day: "", Time: "09:00", Info: "תפילת שחרית מיוחדת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialMincha1", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Day: "", Time: "16:00", Info: "תפילת מנחה מיוחדת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialMincha2", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Day: "", Time: "17:00", Info: "תפילת מנחה מיוחדת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Day: "", Time: "21:00", Info: "תפילת ערבית מיוחדת", On: false})
	c.Default = append(c.Default, ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Day: "", Time: "222:00", Info: "תפילת ערבית מיוחדת", On: false})

	// Marshal, Create and Save
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalln("Can 't Marshak Configuration Default file")
	}

	f, err := os.Create(DEFAULTCONFIGFILE)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(b))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
