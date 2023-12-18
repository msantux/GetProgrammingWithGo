package main

import "fmt"

func main() {
	const distance = 236_000_000_000_000_000
	const lightSpeed = 299_792
	const secondsPerDay = 86_400
	const daysPerYear = 365

	const years = distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("Canis Major Dwarf is", years, "light years away.")
}
