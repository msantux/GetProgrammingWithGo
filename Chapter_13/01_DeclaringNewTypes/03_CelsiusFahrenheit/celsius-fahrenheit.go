package main

import "fmt"

func main() {
	type celsius float64
	type fahrenheit float64

	var c celsius = 20
	var f fahrenheit = 20

	// if c == f {} Invalid Operation: Type mismatch

	// c += f Invelid Operation: Type mismatch

	fmt.Println(c)
	fmt.Println(f)
}
