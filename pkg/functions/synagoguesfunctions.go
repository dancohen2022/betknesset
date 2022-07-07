package functions

import (
	"errors"
	"fmt"

	"github.com/dancohen2022/betknesset/pkg/mdb"
	"github.com/dancohen2022/betknesset/pkg/synagogues"
)

func ResetSynagogueSchedule(name string) bool {
	fmt.Println("ResetSynagogueSchedule")
	sList := mdb.GetAllSynagogues()
	for _, s := range sList {
		syn := s
		if s.User.Name == name {
			mdb.DeleteAllSynagogueSchedules(syn.User.Name)
			//Update API period and other constants
			calend := UpdateApiParamsPeriod(syn.CalendarApi)
			zman := UpdateApiParamsPeriod(syn.ZmanimApi)

			UpdateCalendarJSON(name, calend)
			UpdateZmanimJSON(name, zman)
			UpdateDefaultConfigItemsList(name)

			return true
		}
	}
	return false
}

func UpdateSynagogueSchedule(name string) bool {
	fmt.Println("UpdateSynagogueSchedule")
	sList := mdb.GetAllSynagogues()
	for _, s := range sList {
		syn := s
		if s.User.Name == name {
			mdb.DeleteAllSynagogueSchedules(syn.User.Name)
			//Update API period and other constants
			calend := UpdateApiParamsPeriod(syn.CalendarApi)
			zman := UpdateApiParamsPeriod(syn.ZmanimApi)

			UpdateCalendarJSON(name, calend)
			UpdateZmanimJSON(name, zman)

			return true
		}
	}
	return false
}

func SynagogueExist(name, key string) (synagogues.Synagogue, error) {
	// Used for User Login  with Name and Key
	fmt.Println("SynagogueExist")
	syn := mdb.GetAllSynagogues()
	b := synagogues.Synagogue{}
	for _, s := range syn {
		if (s.User.Name == name) && (s.User.Key == key) {
			b.User.Key = s.User.Key
			b.User.Name = s.User.Name
			b.CalendarApi = s.CalendarApi
			b.ZmanimApi = s.ZmanimApi
			return b, nil
		}
	}
	return b, errors.New("Synagogue doesn't exist")
}

func GetDefaultConfigItemsList() []synagogues.ConfigItem {
	fmt.Println("GetDefaultConfigItemsList")
	/*ConfigItem
	Name     string `json:"name"`
	Hname    string `json:"hname"`
	Category string `json:"category"`
	Subcat   string `json:"subcategory"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Info     string `json:"info"`
	On       bool   `json:"on"`
	*/
	cIList := []synagogues.ConfigItem{}
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekChacharit1", Hname: "תפילת שחרית 1", Category: "tfilot", Date: "", Time: "07:00", Info: "תפילת שחרית", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekChacharit2", Hname: "תפילת שחרית 2", Category: "tfilot", Date: "", Time: "06:00", Info: "תפילת שחרית", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekMincha1", Hname: "תפילת מנחה 1", Category: "tfilot", Date: "", Time: "19:00", Info: "תפילת מנחה", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekMincha2", Hname: "תפילת מנחה 2", Category: "tfilot", Date: "", Time: "15:00", Info: "תפילת מנחה", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekArvit1", Hname: "תפילת ערבית 1", Category: "tfilot", Date: "", Time: "20:00", Info: "תפילת ערבית", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "WeekArvit2", Hname: "תפילת ערבית 2", Category: "tfilot", Date: "", Time: "21:00", Info: "תפילת ערבית", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "ErevShabbatMincha", Hname: "תפילת מנחה ערב שבת", Category: "tfilot", Date: "", Time: "18:00", Info: "תפילת מנחה ערב שבת", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "ShabbatArvit1", Hname: "תפילת ערבית של שבת", Category: "tfilot", Date: "", Time: "18:00", Info: "תפילת ערבית של שבת", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "ShabbatChacharit", Hname: "תפילת שחרית של שבת", Category: "tfilot", Date: "", Time: "08:00", Info: "תפילת שחרית של שבת", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "ShabbatMinha1", Hname: "תפילת מנחה של שבת", Category: "tfilot", Date: "", Time: "17:00", Info: "תפילת מנחה של שבת", On: true})
	cIList = append(cIList, synagogues.ConfigItem{Name: "ShabbatMinha2", Hname: "תפילת מנחה של שבת", Category: "tfilot", Date: "", Time: "14:00", Info: "תפילת מנחה של שבת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialChacharit1", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Date: "", Time: "08:00", Info: "תפילת שחרית מיוחדת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialChacharit2", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Date: "", Time: "09:00", Info: "תפילת שחרית מיוחדת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialMincha1", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Date: "", Time: "16:00", Info: "תפילת מנחה מיוחדת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialMincha2", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Date: "", Time: "17:00", Info: "תפילת מנחה מיוחדת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Date: "", Time: "21:00", Info: "תפילת ערבית מיוחדת", On: false})
	cIList = append(cIList, synagogues.ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Date: "", Time: "222:00", Info: "תפילת ערבית מיוחדת", On: false})

	return cIList

}
