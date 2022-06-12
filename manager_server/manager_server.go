package main

import (
	"fmt"

	"github.com/dancohen2022/betknesset/models"
)

func main() {
	models.InitSynagogues()
	//LoopManager()
	models.CreatFirstDefaultConfigValuesFile()
}

func LoopManager() {

	var name string
	for {
		fmt.Printf("Enter Synagogue name or Enter to exit: ")
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello my name is %s", name)
		if name == "" {
			break
		} else {
			if models.ResetSynagogueSchedule(name) {
				fmt.Printf("%s has been reseted", name)
			} else {
				fmt.Printf("%s doesn't exist", name)
			}
		}
	}
	/*
		for _, s := range *models.GetSynagogues() {
			GetDaySynagogueScheduleJSON(s)
		}
	*/

}
