package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	red = 255
	red += 2
	fmt.Println(red)

	number = 127
	number += 3
	fmt.Println(number)

	red = 0
	red--
	fmt.Println(red)

	number = -128
	number--
	fmt.Println(number)

	var green uint16 = 65535
	green++
	fmt.Println(green)
}
