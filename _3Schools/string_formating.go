package main

import ( // In Go, strings are UTF-8 encoded.
	"fmt"
	"strings"
	"unicode/utf8"
)

type Letter struct {
	Symbol, English string
}

var letters = []Letter{
	{"Σ", "Sigma"},
}

func mainForma() {
	
	words := []string{"«", "Don't", "Panic", "»"}
	//UTF-8 angular brackets
	fmt.Println("Length: ", lineLength(words))

	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))
	fmt.Println(englishFor("�"))
}

func lineLength(words []string) int {
	total := 0
	for _, word := range words {
		total += utf8.RuneCountInString(word)
	}
	numSpaces := len(words) - 1
	return total + numSpaces
}

func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}
	return "", fmt.Errorf("Unknown Greek Letter: %#v", greek)
}
