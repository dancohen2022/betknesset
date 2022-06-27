package synagogues

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
var Synagogues []Synagogue
