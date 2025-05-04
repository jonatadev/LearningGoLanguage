package main

// https://medium.com/openskill/golang-demistifying-interface-value-and-pointer-receiver-365de07c2293
// https://www.youtube.com/watch?v=pBpFH5Nti5k
// https://www.alexedwards.net/blog/interfaces-explained

/*

Interfaces in Go allow you to define a set of method signatures that any type can implement,
providing a flexible and powerful way to write generic code.


*/

type Printer interface {
	Print()
}

/*
func interfaceSample() {

	Print("Jonatas")

}

func (name string) Print() { // Implementação do metódo Print() da interface Printer
	fmt.Println(name)
}

func PrintPerson(p Printer) {
	p.Print()
} */

int sum = 0
for i := 0; i < 10; i++ 
{
	sum += i
}