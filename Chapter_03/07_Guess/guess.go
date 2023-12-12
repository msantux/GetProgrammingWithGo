package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const myNumber = 55
	var guess int

	for {
		guess = rand.Intn(100) + 1

		if guess == myNumber {
			fmt.Printf("You got it! %v\n", guess)
			break
		} else if guess > myNumber {
			fmt.Printf("%v is too big.\n", guess)
		} else {
			fmt.Printf("%v is too small.\n", guess)
		}
	}
}
