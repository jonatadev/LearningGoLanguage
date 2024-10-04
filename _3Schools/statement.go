package main

import "fmt"

func statementSample() {
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else { // The brackets in the else statement should be like } else {
		fmt.Println("Good evening.")
	}

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}
