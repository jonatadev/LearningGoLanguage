package main

import "fmt"

// The for loop is the only loop available in Go.

func for_lop() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// The range keyword is used to more easily iterate over an array, slice or map.
	// It returns both the index and the value.

	fruit := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruit { // idx = index
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// Tip: To only show the value or the index,
	// you can omit the other output using an underscore (_).

	cars := [3]string{"BMW", "Honda", "Mercedes"}
	for _, val := range cars {
		fmt.Printf("%v\n", val)
	}

	// for as a while loop
	total := 1

	for total < 10 {
		total += total
	}

}
