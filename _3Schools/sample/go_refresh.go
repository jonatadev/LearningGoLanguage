package sample

import "fmt"

var counter_sizer int = 10 // use var declare outside of function
var (
	name string = "Jonatas"
	age  int    = 30
)

const Memory_size int = 64 // use const declare outside of function

func mainss() {
	var num = 10              // ou int num = 10
	for i := 0; i < 10; i++ { // não pode quebrar a linha
		num += i
	}

	//Explicit Declaration
	var aString string = "Hello World"
	var sumTwo = 10 + 20
	fmt.Println(aString) // Hello World
	fmt.Println(sumTwo)  // 30

	//Implicit Declaration
	name := "jonatas"
	fmt.Println(name) // jonatas

	// new and make
	// new allocates memory and returns a pointer to the type
	// make allocates memory and returns a value of the type
	// new is used for slices, maps, and channels
	// make is used for slices, maps, and channels
	// new is used for pointers, structs, and arrays

	string_map := new(map[string]string)  // new allocates memory and returns a pointer to the type
	*string_map = make(map[string]string) // make allocates memory and returns a value of the type
	(*string_map)["name"] = "Jonatas"     // set value in map
	(*string_map)["age"] = "30"           // set value in map

	fmt.Println((*string_map)["name"]) // Jonatas
	fmt.Println((*string_map)["age"])  // 30

	// Call by reference
	var surname *string   // a long hexadecimal number
	fmt.Println(*surname) // <nil>

	value1 := 100
	pointer := &value1   // pointer to value1
	fmt.Println(pointer) // 0xc00000a0b8

	// Arrays
	var floatArray = []float32{10.0, 200.0, 35.64, 78.0, 540.60}

	// Slices == ArrayList in Java
	var colors = []string{"red", "green", "blue"}

	// For Loop
	// we used the blank identifier ( _ ) to ignore the index.

	nums := []int{2, 3, 4}
	total := 0

	for _, num := range nums {
		total += num
	}

	for i := range nums {
		fmt.Println(nums[i])
	}

	// For Loop with Blank Pre/Post Statements
	total2 := 1
	for total2 < 10 {
		total2 += total2
	}

	// Methods
	poodle := Dog{"Poodle", 10, "woff!"}
	poodle.Speak() // woff!
}

type Books struct {
	bTitle  string
	bAuthor string
}

// func, parametros, nome da funcao, tipo de retorno
// variable_name type
func (b Books) getTitle() string {
	return b.bTitle
}

func printBookDetails(b *Books) {
	fmt.Println("Book Title: ", b.getTitle())
	fmt.Println("Book Author: ", b.bAuthor)
}

// Functions
// func function_name(parameter_name type) return_type { ... }

func findMaximum(n1, n2 int) int {
	var maxVal int // variable declaration
	if n1 > n2 {
		maxVal = n1
	} else {
		maxVal = n2
	}
	return maxVal
}

// Returnig more than one value
func swapValues(str1, str2 string) (string, string) {
	return str2, str1
}

// Methods
// In Go, if functions are attached to user-defined custom types, they are then
// referred to as methods.

//Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// Predicate function
func filter(pred func(int) bool, values []int) []int {
	var out []int
	for _, val := range values {
		if pred(val) {
			out = append(out, val)
		}
	}
	return out
}

func isOdd(n int) bool {
	return n%2 == 1
}

// values := []int{22, 85, 36, 94, 50, 67, 17, 18}
// fmt.Println("After Filtering: ", filter(isOdd, values))