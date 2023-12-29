package main

import "fmt"

func main() {

	mySlice := []string{"test"}
	sliceCap := cap(mySlice)
	fmt.Println("Capcity 0: ", sliceCap)

	for i := 1; i <= 10000; i++ {
		mySlice = append(mySlice, fmt.Sprint(i))

		if sliceCap < cap(mySlice) {
			sliceCap = cap(mySlice)
			fmt.Printf("Capacity %v: %v\n", i, sliceCap)
		}
	}

	fmt.Println(mySlice)
}
