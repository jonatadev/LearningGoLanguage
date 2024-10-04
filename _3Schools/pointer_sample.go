package main

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.

func pointer() {

	var p *int
	i := 42
	p = &i

	// The * operator denotes the pointer's underlying value.

	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer p
	fmt.Println(*p) // Unlike C, Go has no pointer arithmetic

	value1 := 42
	var pointer1 = &value1
	fmt.Println("Value of pointer1: ", *pointer1) // acessa o valor	
}

// https://medium.com/@jamal.kaksouri/a-comprehensive-guide-to-pointers-in-go-4acc58eb1f4d
