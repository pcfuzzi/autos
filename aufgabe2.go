package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Geben Sie eine Zeichenkette ein: ")
	fmt.Scanln(&input)

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	if input == reverse(input) {
		fmt.Println("Die Zeichenkette ist ein Palindrom")
	} else {
		fmt.Println("Die Zeichenkette ist kein Palindrom")
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
