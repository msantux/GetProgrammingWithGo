package main

import "fmt"

func main() {
	str := "5"
	var launch bool

	switch str {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println("Error:", str, "is not a valid value.")
	}

	fmt.Println("launch =", launch)
}
