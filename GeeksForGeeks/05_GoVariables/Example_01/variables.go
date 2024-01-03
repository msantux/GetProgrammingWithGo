// Go program to illustrate concept of variable
package main

import "fmt"

func main() {
	// Variable declared and initialized without the explicit type
	var myVariable1 = 20
	var myVariable2 = "GeeksforGeeks"
	var myVariable3 = 34.80

	//Display the value and the type of the variables
	fmt.Printf("The value of myVariable1 is: %d\n", myVariable1)
	fmt.Printf("The type of myVariable1 is: %T\n", myVariable1)
	fmt.Printf("The value of myVariable2 is: %s\n", myVariable2)
	fmt.Printf("The type of myVariable2 is: %T\n", myVariable2)
	fmt.Printf("The value of myVariable3 is: %f\n", myVariable3)
	fmt.Printf("The type of myVariable3 is: %T\n", myVariable3)
}
