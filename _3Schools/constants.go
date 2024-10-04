package main

// If a variable should have a fixed value that cannot be changed, you can use the const keyword.
// The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
// Read Only

// const CONSTNAME type = value

const PI = 3.14 // Typed constants

// Multiple constants can be grouped together into a block for readability
const (
	A int = 1
	B     = 3.14
	C     = "HI!"
)

func constantSample() {

	const SUNDAY = "SUNDAY" // Untyped constants are declared without a type
	// When a constant is declared, it is not possible to change the value later
}

// Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
// Constants can be declared both inside and outside of a function
