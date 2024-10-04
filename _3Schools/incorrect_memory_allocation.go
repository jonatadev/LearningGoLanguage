package main

import "fmt"

func mainStruct() {

	//  you are attempting to place data in a map that does not have any
	//  memory storage initialized
	string_Map := new(map[string]int)
	string_Map["Marks"] = 56 // invalid operation: cannot index string_Map (variable of type *map[string]int)
	fmt.Println(string_Map)

	// The correct way to allocate memory for a complex object like the map is to
	// warp your declaration in the make function
	// make declara e aloca memoria
	string_Mapp := make(map[string]int)
	string_Mapp["Marks"] = 56
	fmt.Println(string_Mapp)

	// The make function will also initialize non-zero storage memory for
	// the object.
}
