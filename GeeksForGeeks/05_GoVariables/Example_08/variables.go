// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	// Using short variable declaration
	// Multiple variables of different types are declared and initialized in the single line
	myVar1, myVar2, myVar3 := 800, "Geeks", 47.56

	//Display the value and the type of the variables
	fmt.Printf("The value of myVar1 is: %d\n", myVar1)
	fmt.Printf("The type of myVar1 is: %T\n\n", myVar1)
	fmt.Printf("The value of myVar2 is: %s\n", myVar2)
	fmt.Printf("The type of myVar2 is: %T\n\n", myVar2)
	fmt.Printf("The value of myVar3 is: %f\n", myVar3)
	fmt.Printf("The type of myVar3 is: %T\n\n", myVar3)
}
