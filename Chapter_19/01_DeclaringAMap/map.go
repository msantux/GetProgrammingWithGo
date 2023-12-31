package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vºC.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	if moon2, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vºC.\n", moon2)
	} else {
		fmt.Println("Where is the moon?")
	}
}
