package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func mainIf() {

	// An if Statement with Initialization

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than Zero"
	} else if theAnswer == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than Zero"
	}
	fmt.Println(result)
}
