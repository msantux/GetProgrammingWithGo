package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	//Quick check 2.6 (page 21)
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println("Distance to Mars =", distance)
}
