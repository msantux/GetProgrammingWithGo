package main

import "fmt"

func main() {
	var room = "lake"

	switch room {
	case "cave":
		fmt.Println("You find yoursel in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}
}
