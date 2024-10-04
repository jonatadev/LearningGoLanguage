package main

import "fmt"

// Maps are used to store data values in key:value pairs.
// Each element in a map is a key:value pair.
// Each element in a map is a key:value pair.
// The default value of a map is nil.

func main_() {

	// The order of the map elements defined in the code is different from the way that they are stored.
	var country = map[string]string{"Brazil": "Brasilia", "UK": "London"}
	fmt.Printf("country\t%v\n", country)

	var car = make(map[string]string) // the map is empty now
	car["brand"] = "Ford"
	car["model"] = "Mustang"
	car["year"] = "1964"

	// Iterate Over Maps
	// You can use range to iterate over maps.

	for k, v := range car {
		fmt.Printf("%v : %v, ", k, v)
	}

	//for key, value := range map_variable_name{}

}

// Note that key and value are both user-defined variable names.
