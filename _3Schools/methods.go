package main

// In Go a method is a member of a type.
// Go doesn’t support method overriding; each method must have its
// own unique name.

import (
	"fmt"
)

func mainMeth() {

	poodle := Dog{"Poodle", 10, "woff!"}

	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v \n", poodle.Breed, poodle.Weight)

	poodle.Speak()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}
