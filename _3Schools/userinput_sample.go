package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func mainBufio() {

	var (
	//name string
	//age  int
	)

	//var name string
	//fmt.Print("Enter your name: ")
	//fmt.Scanf("%s", name)
	//fmt.Println("Hello ", name)

	fmt.Println("PI = ", math.Pi)
	now := time.Now()
	fmt.Println("Now = ", now)

	// Using bufio
	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Some Text: ")
	// func (b *bufio.Reader) ReadString(delim byte) (string, error)
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered: ", input)

} // if you want to ignore a variable in Go, you name it with an underscore.
