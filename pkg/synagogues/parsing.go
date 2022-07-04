package synagogues

type Item interface {
	ParseItem() []ConfigItem
}

func ParseItemsList(items []Item) []ConfigItem {
	lst := items
	lst_daily := []ConfigItem{}
	for _, v := range lst {
		d := v.ParseItem()
		lst_daily = append(lst_daily, d...)
	}
	return lst_daily
}

func (item CalendarItems) ParseItem() []ConfigItem {
	var d ConfigItem
	d.Category = item.Category
	d.Date = item.Date
	d.Hname = item.Hebrew
	d.Info = item.Memo
	d.Name = item.Title
	d.Subcat = item.Subcat
	d.Time = ""
	d.On = true
	return append([]ConfigItem{}, d)
}

func (item ZmanimJson) ParseItem() []ConfigItem {
	lst := []ConfigItem{}
	timesMap := GetItemsTimeMap(item.Times)

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
			lst = append(lst, d)
		}
	}

	return lst
}

func GetItemsTimeMap(Times ZmanimTimes) map[string]map[string]string {
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
