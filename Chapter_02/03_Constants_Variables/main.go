// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   //km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 96300000
	var speed int = 100800 // km/h
	const hoursPerDay = 24
	fmt.Println((distance/speed)/hoursPerDay, "days")

}
