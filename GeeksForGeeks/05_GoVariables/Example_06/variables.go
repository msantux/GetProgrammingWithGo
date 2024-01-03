// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	//Using short variable declaration
	// Multiple variables of same types are declared and initialized in the single line
	myVar1, myVar2, myVar3 := 800, 34, 56

	//Display the value and the type of the variables
	fmt.Printf("The value of myVar1 is: %d\n", myVar1)
	fmt.Printf("The type of myVar1 is: %T\n", myVar1)
	fmt.Printf("\nThe value of myVar2 is: %d\n", myVar2)
	fmt.Printf("The type of myVar2 is: %T\n", myVar2)
	fmt.Printf("\nThe value of myVar3 is: %d\n", myVar3)
	fmt.Printf("The type of myVar3 is: %T\n", myVar3)
}
