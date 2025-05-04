package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var balanceAmount int
var mutex = &sync.Mutex{}

func creditToBalance() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balanceAmount += 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Credit: ",
			balanceAmount)
		mutex.Unlock()
	}
}

func debitFromBalance() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balanceAmount -= 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Debit: ",
			balanceAmount)
		mutex.Unlock()
	}
}

// In this code, Lock() and Unlock() are built-in functions that can be
// used with a mutex type object and can mark the start and end of a critical
// section of the program.

func main() {
	balanceAmount = 3000
	fmt.Println("Initial balance: ", balanceAmount)
	go creditToBalance()
	go debitFromBalance()
	fmt.Scanln()
}
