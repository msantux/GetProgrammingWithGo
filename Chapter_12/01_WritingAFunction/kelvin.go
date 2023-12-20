package main

import "fmt"

// kelvinToCelsius converts K to ºC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)

	fmt.Println(kelvin, "K is", celsius, "ºC")
}
