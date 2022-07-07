package synagogues

import (
	"fmt"
	"strings"
	"time"
)

func ParseCalendarItemsToConfigItems(cList []CalendarItems) []ConfigItem {
	fmt.Println("ParseCalendarItemsToConfigItems")
	newList := []ConfigItem{}
	for _, item := range cList {
		var c ConfigItem
		c.Category = item.Category
		c.Date = DateFormat(item.Date)
		c.Hname = item.Hebrew
		c.Info = item.Memo
		c.Name = item.Title
		c.Subcat = item.Subcat
		c.Time = ""
		c.On = true

		fmt.Printf("c: %v\n", c)
		fmt.Printf("c.Date: %v\n", c.Date)
		fmt.Printf("c.Name: %v\n", c.Name)
		fmt.Printf("c.Hname: %v\n", c.Hname)
		fmt.Printf("c.Info: %v\n", c.Info)
		fmt.Printf("c.Category: %v\n", c.Category)
		fmt.Printf("c.Subcat: %v\n", c.Subcat)
		fmt.Printf("c.Time: %v\n", c.Time)
		fmt.Printf("c.On: %v\n", c.On)

		newList = append(newList, c)
	}
	//fmt.Printf("newList: %v\n", newList)
	return newList
}

func ParseZmanimJsonToConfigItems(zm ZmanimJson) []ConfigItem {
	fmt.Println("ParseZmanimJsonToConfigItems")
	newList := []ConfigItem{}
	timesMap := GetItemsTimeMap(zm.Times)

	for key1, val1 := range timesMap {
		name := key1
		for key2, val2 := range val1 {
			k := key2
			v := val2
			var d ConfigItem
			d.Category = "dailyTimes"
			d.Date = k
			d.Hname = SetHebrewName(name)
			d.Info = ""
			d.Name = name
			d.Subcat = ""
			d.Time = v
			d.On = true
			fmt.Printf("d: %v\n", d)
			fmt.Printf("d.Name: %v\n", d.Name)
			fmt.Printf("d.Date: %v\n", d.Date)
			fmt.Printf("d.Hname: %v\n", d.Hname)
			fmt.Printf("d.Info: %v\n", d.Info)
			fmt.Printf("d.Category: %v\n", d.Category)
			fmt.Printf("d.Subcat: %v\n", d.Subcat)
			fmt.Printf("d.Time: %v\n", d.Time)
			fmt.Printf("d.On: %v\n", d.On)
			newList = append(newList, d)
		}
	}
	//fmt.Printf("newList: %v\n", newList)

	return newList
}

func GetItemsTimeMap(Times ZmanimTimes) map[string]map[string]string {
	fmt.Println("GetItemsTimeMap")
	timesMap := make(map[string]map[string]string)
	timesMap["ChatzotNight"] = Times.ChatzotNight
	timesMap["AlotHaShachar"] = Times.AlotHaShachar
	timesMap["Misheyakir"] = Times.Misheyakir
	timesMap["MisheyakirMachmir"] = Times.MisheyakirMachmir
	timesMap["Dawn"] = Times.Dawn
	timesMap["Sunrise"] = Times.Sunrise
	timesMap["SofZmanShma"] = Times.SofZmanShma
	timesMap["SofZmanShmaMGA"] = Times.SofZmanShmaMGA
	timesMap["SofZmanTfilla"] = Times.SofZmanTfilla
	timesMap["SofZmanTfillaMGA"] = Times.SofZmanTfillaMGA
	timesMap["Chatzot"] = Times.Chatzot
	timesMap["MinchaGedola"] = Times.MinchaGedola
	timesMap["MinchaKetana"] = Times.MinchaKetana
	timesMap["PlagHaMincha"] = Times.PlagHaMincha
	timesMap["Sunset"] = Times.Sunset
	timesMap["Dusk"] = Times.Dusk
	timesMap["Tzeit7083deg"] = Times.Tzeit7083deg
	timesMap["Tzeit85deg"] = Times.Tzeit85deg
	timesMap["Tzeit42min"] = Times.Tzeit42min
	timesMap["Tzeit50min"] = Times.Tzeit50min
	timesMap["Tzeit72min"] = Times.Tzeit72min

	return timesMap
}

func SetHebrewName(name string) string {
	fmt.Println("SetHebrewName")
	switch name {
	case "ChatzotNight":
		return "חצות הלילה"
	case "AlotHaShachar":
		return "עלות השחר"
	case "Misheyakir":
		return "מישיקיר"
	case "MisheyakirMachmir":
		return "מישיקיר משמיר"
	case "Dawn":
		return "שחר"
	case "Sunrise":
		return "זריחה"
	case "SofZmanShma":
		return "סוף זמם קריאת שמע"
	case "SofZmanShmaMGA":
		return "סוף זמם קריאת שמע גרא"
	case "SofZmanTfilla":
		return "סוף זמן תפילה"
	case "SofZmanTfillaMGA":
		return "סוף זמן תפילה גרא"
	case "Chatzot":
		return "חצות היום"
	case "MinchaGedola":
		return "מנחה גדולה"
	case "MinchaKetana":
		return "מנחה קטנה"
	case "PlagHaMincha":
		return "פלג מנחה"
	case "Sunset":
		return "שקיעה"
	case "Dusk":
		return "אפלולית"
	case "Tzeit7083deg":
		return "צאת הכוכבים 7.083"
	case "Tzeit85deg":
		return "צאת הכוכבים 8.5"
	case "Tzeit42min":
		return "צאת הכוכבים אחרי 42 דקות"
	case "Tzeit50min":
		return "צאת הכוכבים אחרי 50 דקות"
	case "Tzeit72min":
		return "צאת הכוכבים אחרי 72 דקות"
	default:
		return "זמנים"
	}
}

const PERIOD int = 3 //7

func UpdateParamsPeriod(api string) string {
	fmt.Println("UpdateApiParams")
	fmt.Printf("api: %s \n", api)

	//Files limited period

	bStart := DateFormat(time.Now().String())
	bEnd := DateFormat(time.Now().AddDate(0, 0, PERIOD).String())
	periodStart := strings.TrimSpace(fmt.Sprintf("start=%s", bStart))
	periodEnd := fmt.Sprintf("end=%s", bEnd)
	newPeriod := fmt.Sprintf("&%s&%s", periodStart, periodEnd)
	newApi := strings.Replace(api, "&year=now", "", 3)
	newApiTrimes := strings.TrimSpace(newApi + newPeriod)
	fmt.Printf("newApiTrimes: %s \n", newApiTrimes)

	return newApiTrimes

}

func DateFormat(dateString string) string {
	fmt.Println("DateFormat")
	//YYYY-MM-DD
	d := []byte(dateString)
	d = d[:11]
	s := string(d)
	return s
}
