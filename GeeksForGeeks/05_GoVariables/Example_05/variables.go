// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	//Using short variable declaration
	myVar1 := 39
	myVar2 := "GeeksforGeeks"
	myVar3 := 34.67

	//Display the value and the type of the variables
	fmt.Printf("The value of myVar1 is: %d\n", myVar1)
	fmt.Printf("The type of myVar1 is: %T\n", myVar1)
	fmt.Printf("\nThe value of myVar2 is: %s\n", myVar2)
	fmt.Printf("The type of myVar2 is: %T\n", myVar2)
	fmt.Printf("\nThe value of myVar3 is: %f\n", myVar3)
	fmt.Printf("The type of myVar3 is: %T\n", myVar3)
}
