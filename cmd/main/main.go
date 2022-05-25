package main

import (
	"fmt"

	"github.com/dancohen2022/betknesset/pkg/models"
)

func main() {
	// Get synagogs list
	models.InitSynagogues()
	//

	fmt.Println("End Main")
}
