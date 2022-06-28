package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dancohen2022/betknesset/pkg/synagogues"
)

func InitSynagogues() {

	bFile2, err := ioutil.ReadFile(synagogues.SYNAGOGUESFILE)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bFile2, &synagogues.Synagogues)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetSynagogues() *[]synagogues.Synagogue {
	fmt.Printf("synagogues: %v\n\n\n", synagogues.Synagogues)
	return &synagogues.Synagogues
}

func ResetSynagogueSchedule(name string) bool {
	for _, s := range synagogues.Synagogues {
		if s.User.Name == name {
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

func SynagogueExist(name, key string) (synagogues.Synagogue, error) {
	// Used for User Login  with Name and Key
	fmt.Println("SynagogueExist")
	syn := synagogues.Synagogues
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

func InitSSynagogueConfig(synagogue synagogues.Synagogue) *synagogues.ConfigJson {

	var config synagogues.ConfigJson
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

func SetConfigDefValues(synagogues.ConfigJson) *[]synagogues.ConfigItem {
	s, err := ioutil.ReadFile(synagogues.DEFAULTCONFIGFILE)
	c := []synagogues.ConfigItem{}
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
	c := synagogues.ConfigJson{}
	c.Info.Name = "Synagogue Name"
	c.Info.Hname = "שם בית כנסת"
	c.Info.Logo = "logo"
	c.Info.Background = "background"
	c.Info.Message = "about the synagogue"
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekChacharit1", Hname: "תפילת שחרית 1", Category: "tfilot", Date: "", Time: "07:00", Info: "תפילת שחרית", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekChacharit2", Hname: "תפילת שחרית 2", Category: "tfilot", Date: "", Time: "06:00", Info: "תפילת שחרית", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekMincha1", Hname: "תפילת מנחה 1", Category: "tfilot", Date: "", Time: "19:00", Info: "תפילת מנחה", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekMincha2", Hname: "תפילת מנחה 2", Category: "tfilot", Date: "", Time: "15:00", Info: "תפילת מנחה", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekArvit1", Hname: "תפילת ערבית 1", Category: "tfilot", Date: "", Time: "20:00", Info: "תפילת ערבית", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "WeekArvit2", Hname: "תפילת ערבית 2", Category: "tfilot", Date: "", Time: "21:00", Info: "תפילת ערבית", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "ErevShabbatMincha", Hname: "תפילת מנחה ערב שבת", Category: "tfilot", Date: "", Time: "18:00", Info: "תפילת מנחה ערב שבת", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "ShabbatArvit1", Hname: "תפילת ערבית של שבת", Category: "tfilot", Date: "", Time: "18:00", Info: "תפילת ערבית של שבת", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "ShabbatChacharit", Hname: "תפילת שחרית של שבת", Category: "tfilot", Date: "", Time: "08:00", Info: "תפילת שחרית של שבת", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "ShabbatMinha1", Hname: "תפילת מנחה של שבת", Category: "tfilot", Date: "", Time: "17:00", Info: "תפילת מנחה של שבת", On: true})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "ShabbatMinha2", Hname: "תפילת מנחה של שבת", Category: "tfilot", Date: "", Time: "14:00", Info: "תפילת מנחה של שבת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialChacharit1", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Date: "", Time: "08:00", Info: "תפילת שחרית מיוחדת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialChacharit2", Hname: "תפילת שחרית מיוחדת", Category: "tfilot", Date: "", Time: "09:00", Info: "תפילת שחרית מיוחדת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialMincha1", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Date: "", Time: "16:00", Info: "תפילת מנחה מיוחדת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialMincha2", Hname: "תפילת מנחה מיוחדת", Category: "tfilot", Date: "", Time: "17:00", Info: "תפילת מנחה מיוחדת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Date: "", Time: "21:00", Info: "תפילת ערבית מיוחדת", On: false})
	c.Default = append(c.Default, synagogues.ConfigItem{Name: "SpecialArvit1", Hname: "תפילת ערבית מיוחדת", Category: "tfilot", Date: "", Time: "222:00", Info: "תפילת ערבית מיוחדת", On: false})

	// Marshal, Create and Save
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalln("Can 't Marshak Configuration Default file")
	}

	f, err := os.Create(synagogues.DEFAULTCONFIGFILE)

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
