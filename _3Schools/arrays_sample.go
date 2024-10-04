package main

import "fmt"

// Arrays are used to store multiple values of the same type in a single variable

func arraySample() {

	var i int
	var fruits = [3]string{"apple", "lemon", "mango"}
	cars := [...]string{"Bmw", "Porche"}

	fmt.Println("Elements stored in fruits: ")

	for i = 0; i < 3; i++ {
		fmt.Println("Element[%d] = %f \n", i, fruits[i])
	}

	fmt.Println("Size of fruits = ", len(fruits))

	fmt.Println(fruits)
	fmt.Println(cars)

	prices := [3]int{10, 20, 30}
	fmt.Println(prices)
	prices[2] = 50

	fmt.Println(len(prices)) // The len() function is used to find the length of an array
}
