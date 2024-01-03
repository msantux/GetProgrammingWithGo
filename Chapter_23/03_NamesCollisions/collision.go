package main

import "fmt"

type sol int

type celsius float64

type location struct {
	lat, long float64
}

type temperature struct {
	high, low celsius
}

type report struct {
	sol
	temperature
	location
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	// To-do: complicated distance calculation
	return 5
}

func main() {

	report := report{
		sol:         15,
		location:    location{lat: -4.5895, long: 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("%+v\n", report)
	fmt.Printf("average %vºC\n", report.average())

	fmt.Printf("%vºC\n", report.high)
	report.high = 32
	fmt.Printf("%vºC\n", report.temperature.high)

	fmt.Println(report.sol.days(1446))
	// fmt.Println(report.days(1446))
}
