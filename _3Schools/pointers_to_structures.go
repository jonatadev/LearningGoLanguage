package main

// Pointers to Structures

import (
	"fmt"
)

type Book struct {
	bTitle      string
	bAuthorName string
	bSubject    string
	book_id     int
}

func mainP() {

	var Book1 Book
	Book1.bTitle = "The Go Programming Language"
	Book1.bAuthorName = "Alan A. A. Donovan and Brian W. Kernighan"
	Book1.bSubject = "A complete guide to Go programming"
	Book1.book_id = 6495

	/* Print the details of Book1 by passing it as a pointer to function */

	printBookDetails(&Book1) // endereço do ponteiro (referência)

}

func printBookDetails(book *Book) { // passando ponteiro
	fmt.Printf("\nTitle : %s\n", book.bTitle)
	fmt.Printf("Authors : %s\n", book.bAuthorName)
	fmt.Printf("Subject : %s\n", book.bSubject)
	fmt.Printf("Book ID : %d\n", book.book_id)
}

/*

Structs can encapsulate data and methods.

The struct type in Go is a data structure that has similar purpose
and function as the struct type in C and as classes in Java.

Each structure is independent with its own fields for data management
and optionally its own methods.




*/
