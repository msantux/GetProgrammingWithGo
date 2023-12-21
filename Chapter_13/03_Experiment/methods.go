package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * (5.0 / 9.0))
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin(f.celsius() + 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k.celsius()*(9.0/5.0) + 32.0)
}

func main() {
	var c celsius = 0

	fmt.Print(c, "ºC is ", c.kelvin(), "K\n")
	fmt.Print(c, "ºC is ", c.fahrenheit(), "ºF\n")

	var f fahrenheit = 0
	fmt.Print(f, "ºF is ", f.celsius(), "ºC\n")
	fmt.Print(f, "ºF is ", f.kelvin(), "K\n")

	var k kelvin = 0
	fmt.Print(k, "K is ", k.celsius(), "ºC\n")
	fmt.Print(k, "K is ", k.fahrenheit(), "ºF\n")

}
