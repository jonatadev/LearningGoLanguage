package main

import (
	"fmt"
	"regexp"
)

func main() {

	/* Checking directly if the regex pattern matches a string */
	match, _ := regexp.MatchString("p([a-z]+)ch", "preach")
	fmt.Println(match)
}
