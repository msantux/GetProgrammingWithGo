package main

import "fmt"

type celsius float64

type location struct {
	lat, long float64
}

type temperature struct {
	high, low celsius
}

type report struct {
	sol         int
	temperature temperature
	location    location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{lat: -4.5895, long: 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %vºC\n", report.temperature.high)
	fmt.Printf("average %vºC\n", report.temperature.average())
}
