package main

import "fmt"

func main() {
	answer := 42
	address := &answer

	fmt.Printf("address is a %T\n", address)

	canada := "Canada"

	var home *string
	fmt.Printf("home is %T\n", home)

	home = &canada
	fmt.Println(*home)
}
