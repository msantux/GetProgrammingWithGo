package main

import "fmt"

// kelvinToCelsius converts K to ºC
func kelvinToCelsius(temperature float64) float64 {
	temperature -= 273.15
	return temperature
}

// celsiusToFahrenheit converts ºC to ºF
func celsiusToFahrenheit(temperature float64) float64 {
	temperature = (temperature * 9.0 / 5.0) + 32.0
	return temperature
}

// kelvinToFahrenheit converts K to ºF
func kelvinToFahrenheit(temperature float64) float64 {
	temperature = kelvinToCelsius(temperature)
	temperature = celsiusToFahrenheit(temperature)

	return temperature
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)

	fmt.Println(kelvin, "K is", celsius, "ºC")

	celsius = 0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println(celsius, "ºC is", fahrenheit, "ºF")

	kelvin = 0
	fahrenheit = kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "K is", fahrenheit, "ºF")
}
