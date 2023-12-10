package main

import "fmt"

func main() {
	var speed float64       // km/h
	var distance = 56000000 //km
	var travelDuration = 28 //days

	const hoursPerDay = 24

	speed = float64(distance) / float64(travelDuration*hoursPerDay)
	fmt.Printf("Speed to reach Malacandra in 28 days = %.2f km/h", speed)
}
