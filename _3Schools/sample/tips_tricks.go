package sample

import (
	"fmt"
	jonatas_fmt "fmt" // import alias_name "package_name"
	. "regex"
)

// import . "package_name"

func mainS() {

	jonatas_fmt.Println("Hello World")
	//text := .MatchString("m([a-z]+)rs", "mars")
	//jonatas_fmt.Println(text)

	user := CreateNewUser(10, "Maryam", "PK")
	fmt.Println(user.GreetUser())
}

type UserData struct {
	ID       int
	Name     string
	Location string
}

func CreateNewUser(id int, name, location string) *UserData {
	id++
	return &UserData{id, name, location}
}

func (user *UserData) GreetUser() string {
	return fmt.Sprintf("Hello %s from %s",
		user.Name, user.Location)
}