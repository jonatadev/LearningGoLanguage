package main

import (
	"fmt"
	"os"
)

func mainDefer() {

	f := createFile("defer-eg.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(fileName string) *os.File {
	fmt.Println("Creating File with name ", fileName)
	file, err := os.Create(fileName)
	if err != nil { // In Go, the panic idiom is used to return errors.
		panic(err)
	}
	return file
}

func writeFile(fileName *os.File) {
	fmt.Println("Writing Data to File with name ",
		fileName.Name())
	fmt.Fprintln(fileName, "Hi, My name is Maryam.")
}

func closeFile(file *os.File) {
	fmt.Println("Closing File with name ", file.Name())
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
