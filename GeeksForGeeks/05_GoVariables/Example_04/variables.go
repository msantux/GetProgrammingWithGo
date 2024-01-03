// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	// Multiple variable of different types are declared and initialized in the single line
	var myVariable1, myVariable2, myVariable3 = 2, "GFG", 67.56

	//Display the value and the type of the variables
	fmt.Printf("The value of myVariable1 is: %d\n", myVariable1)
	fmt.Printf("The type of myVariable1 is: %T\n", myVariable1)
	fmt.Printf("The value of myVariable2 is: %s\n", myVariable2)
	fmt.Printf("The type of myVariable2 is: %T\n", myVariable2)
	fmt.Printf("The value of myVariable3 is: %f\n", myVariable3)
	fmt.Printf("The type of myVariable3 is: %T\n", myVariable3)
}
