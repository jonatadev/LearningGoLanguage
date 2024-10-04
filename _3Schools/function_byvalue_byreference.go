package main

import (
	"fmt"
	//"io"
)

func mainPass() {

	num := 500
	fmt.Println("Value of Passed Argument Before Function Call: ", num)
	fmt.Println("Value Returned by Function call:", modifyInt(num))
	fmt.Println("Value of Passed Argument After Function Call: ", num)
}

func modifyInt(val int) int {
	return val + 5
}
