package main

// A function is a block of statements that can be used repeatedly in a program.
// Use the func keyword.
// Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.
// Function names are case-sensitive
// Give the function a name that reflects what the function does!

import "fmt"


func mainSample() {

	fmt.Println(add(43, 7))
	familyName("Jonatas Alves da Silva", 40)

	// Store the return value in a variable
	total := sumV(20, 80)
	fmt.Println(total)

	// multiple return value
	// fmt.Println(multipleReturn(5, "Hello"))

	// Here, we want to omit the first returned value
	// (result - which is stored in variable a)
	_, b := multipleReturn(5, "Hello")
	fmt.Println(b) // Hello World!

	val1, val2 := swapValues("Abdullah", "Hassan")
    fmt.Println("Values After Swap: ", val1, val2)

}

// function with parameter
func familyName(first_name string, age int) {
	fmt.Println(first_name, age)
}

// Retunr value
// func FunctionName(param1 type, param2 type) type {
func add(x int, y int) int {
	return x + y
}

// named retunr values
// In Go, you can name the return values of a function.

func sumV(value_a int, value_b int) (total int) {
	total = value_a + value_b
	return // naked return, retorna sem especificar o nome de uma variável so return
}

// Go functions can also return multiple values.
func multipleReturn(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func addValues(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func swapValues(str1, str2 string) (string, string) {
	return str2, str1
}
