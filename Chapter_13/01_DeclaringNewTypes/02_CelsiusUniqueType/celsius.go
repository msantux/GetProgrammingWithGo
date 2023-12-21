package main

import "fmt"

func main() {
	type celsius float64

	var temperature celsius = 20
	var warmUp float64 = 10

	temperature += celsius(warmUp)

	fmt.Println(temperature)
}
