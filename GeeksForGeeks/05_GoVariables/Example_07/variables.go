// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	//Using short variable declaration
	// Here, short variable declaration acts as an assignment for myVar2 variable
	// because same variable present in the same block so the value of myVar2 is
	// changed from 45 to 100
	myVar1, myVar2 := 39, 45
	myVar3, myVar2 := 45, 100

	// If you try to run the commented lines, then compiler will gives error because
	// these variables are already defined
	// myVar1, myVar2 := 43, 47
	// myVar3 := 200

	//Display the value and the type of the variables
	fmt.Printf("The value of myVar1 is: %d\n", myVar1)
	fmt.Printf("The value of myVar2 is: %d\n", myVar2)
	fmt.Printf("The value of myVar3 is: %d\n", myVar3)
}
