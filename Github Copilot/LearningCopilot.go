// GitHub Copilot is an AI pair programmer tool that helps you write code faster and smarter.
// https://code.visualstudio.com/docs/copilot/overview

// https://marketplace.visualstudio.com/items?itemName=GitHub.copilot
// https://code.visualstudio.com/docs/copilot/getting-started
// create a code in go to sum elements of an array

// CTRL + I ==> Adiciona Copilot no codigo

package main

import "fmt"

func average(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Average:", average(arr))
}

