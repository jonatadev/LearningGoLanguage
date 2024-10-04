package main

import "fmt"

// Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
// when written correctly recursion can be a very efficient and mathematically-elegant approach to programming.

/*
func main() {
	testCount(1)
	fmt.Println(factorial(4)) // 24
} */

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func factorial(x float64) (y float64) {
	if x > 0 {
		y = x * factorial(x - 1)
	} else {
		y = 1
	}
	return
}
