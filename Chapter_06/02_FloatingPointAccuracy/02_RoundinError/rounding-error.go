package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "º F\n")
	fmt.Print((9.0/5.0*celsius)+32, "º F\n")

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "º F\n")
}
