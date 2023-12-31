// Built-in functions in complex numbers
package main

import "fmt"

func main() {
	comp1 := complex(10, 11)

	// complex number init syntax
	comp2 := 13 + 33i

	fmt.Println("Complex number 1 is: ", comp1)
	fmt.Println("Complex number 2 is: ", comp2)

	// get real part
	realNum := real(comp1)
	fmt.Println("Real part of complex number 1 is: ", realNum)

	//get imaginary part
	imaginary := imag(comp2)
	fmt.Println("Imaginary part of complex number 2 is: ", imaginary)
}
