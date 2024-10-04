package main

import (
	"fmt"
)

func mainCalc() {

	fmt.Println("Mean of {1,2,3} is: ", mean([]int{1, 2, 3}))
	fmt.Println("Mean of {1,2,3,4,5} is: ", mean([]int{1, 2, 3, 4, 5}))

	var moby = []string{"shall", "pass", "shall", "away", "too"}

	fmt.Println(frequency(moby))
}

func mean(nums []int) float64 {
	s := sum(nums)
	return (float64(s) / float64(len(nums)))
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func frequency(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	return freq
}
