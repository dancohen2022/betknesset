package main

import (
	"fmt"

	"github.com/dancohen2022/betknesset/pkg"
)

func main() {
	pkg.InitSynagogues()
	//LoopManager()
	pkg.CreatFirstDefaultConfigValuesFile()
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
			if pkg.ResetSynagogueSchedule(name) {
				fmt.Printf("%s has been reseted", name)
			} else {
				fmt.Printf("%s doesn't exist", name)
			}
		}
	}
	/*
		for _, s := range *pkg.GetSynagogues() {
			GetDaySynagogueScheduleJSON(s)
		}
	*/

}
