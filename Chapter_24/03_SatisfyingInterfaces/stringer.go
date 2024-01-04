package main

import (
	"fmt"
)

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long float64
}

// String formats a location with latitude and longitude.
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

// String formats a DMS coordinate
func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}

// locationsDMS with latitude, longitude in DMS
type locationDMS struct {
	lat, long coordinate
}

// String formats a locationDMS with latitude and longitude.
func (l locationDMS) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	curiosity := location{lat: -4.5895, long: 137.4417}
	fmt.Println(curiosity)

	elysium := locationDMS{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium)
}
