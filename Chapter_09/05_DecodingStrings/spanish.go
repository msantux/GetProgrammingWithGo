package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Comó estás?"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)

	fmt.Printf("First rune: %c %v bytes\n", c, size)

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}
