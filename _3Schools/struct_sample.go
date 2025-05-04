package main

import "fmt"

// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
// structs are used to store multiple values of different data types into a single variable.
// A struct can be useful for grouping data together to create records.

// type struct_name struct {}

func mainStruct() {

	var jonatas Person

	jonatas.name = "Jonatas Alves"
	jonatas.age = 100
	jonatas.email = "dev@gmail.com"

	fmt.Println("Name = ", jonatas.name)
	fmt.Println()
	printPerson(jonatas)

}

type Person struct { // Notice that the struct members above have different data types.
	name  string
	age   int
	email string
}

type Books struct {
	bTitle      string
	bAuthorName string
	bSubject    string
	book_id     int
}

// Pass Struct as Function Arguments

func printPerson(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Email: ", person.email)
}
