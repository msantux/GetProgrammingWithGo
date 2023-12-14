package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%-16v %-4v %-10v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		var spaceline string

		switch selection := rand.Intn(3) + 1; selection {
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "SpaceX"
		default:
			spaceline = "Virgin Galactic"
		}

		const distanceToMars = 62_100_000
		const secondsInADay = 60 * 60 * 24
		var travelSpeed = rand.Intn(15) + 16                                    //km/sec
		var voyageDuration int = (distanceToMars / travelSpeed) / secondsInADay //days

		var price = 20 + travelSpeed

		var trypType string

		if i := rand.Intn(2); i == 0 {
			trypType = "One-way"
		} else {
			trypType = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, voyageDuration, trypType, price)
	}

}
