package main

import (
	"fmt"
	"math/rand"
	"time"
)

var balanceAmount int

func creditToBalance() {
	for i := 0; i < 5; i++ {
		balanceAmount += 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Credit: ",
			balanceAmount)
	}
}

func debitFromBalance() {
	for i := 0; i < 5; i++ {
		balanceAmount -= 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Debit: ",
			balanceAmount)
	}
}

func main() {
	
	balanceAmount = 3000
	fmt.Println("Initial balance: ", balanceAmount)
	go creditToBalance()
	go debitFromBalance()
	fmt.Scanln()
}
