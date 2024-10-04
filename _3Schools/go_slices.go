package main

// https://pkg.go.dev/sort

import (
	"fmt"
	//"sort"
)

//  unlike arrays, the length of a slice can grow and shrink as you see fit.

func sliceSample() {

	// slice_name := []datatype{values}
	brazil := []string{"Rio de janeiro", "São Paulo"}
	fmt.Println(len(brazil))
	fmt.Println(cap(brazil))

	brazil = append(brazil, "Bahia")

	var colors = []string{"ed", "Green", "Blue"}
	var markss []float32            /* slice of unspecified size */
	var marksss = make([]int, 5, 5) /* slice of length and capacity 5*/

	// A slice is initialized as nil by default if it is declared with no inputs.
	var marks []float64

	if marks == nil {
		fmt.Println("Slice is nil")
	}

	printItemsOfSlice(marks) //pass slice to a function

}

func printItemsOfSlice(x []float64) {
	fmt.Printf("Length=%d Capacity=%d Slice=%v\n", len(x), cap(x), x)
}
