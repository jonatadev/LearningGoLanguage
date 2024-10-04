package main

import (
	"fmt"
)

var(
	name string
	email string
	age int
)

var size int
var date string

// In Go, it is possible to declare multiple variables in the same line.
var a, b, c, d int = 1, 3, 5, 7

func variables() {
	// var variablename type = value
	// You always have to specify either type or value (or both)

	var email string = "jonatadev@gmail.com" // In this case, the type of the variable is inferred from the value
	username := "jonatas" // quando não usar o var o simbolo de atribuição é :=

	fmt.Println(username)
	fmt.Println(email)

	// In Go, all variables are initialized. So, if you declare a variable without an initial
	// value, its value will be set to the default value of its type.

	var account_number int      // These variables are declared but they have not been assigned initial values.
	fmt.Println(account_number) // ""

	// Multiple variable declarations can also be grouped together into a block for greater readability:
	var (
		login    string
		password int
		session  int
	)

	fmt.Println(login)
	fmt.Println(password)
	fmt.Println(session)
}

// Use the var keyword, followed by variable name and type:
