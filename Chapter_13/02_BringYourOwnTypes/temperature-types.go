package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts K to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// celsiusToKelvin converts ºC to K
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0

	c := kelvinToCelsius(k)
	fmt.Print(k, "K is ", c, "ºC\n")

	c = 127
	k = celsiusToKelvin(c)

	fmt.Print(c, "ºC is ", k, "K\n")
}
