package main

import "fmt"

func main() {
	launch := false

	lauchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", lauchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	yesNo = "yes"
	launch = (yesNo == "yes")
	fmt.Println("Ready for launch:", launch)
}
