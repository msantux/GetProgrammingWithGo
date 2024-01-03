package main

import "fmt"

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere
type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	//Bradbury Landing 4º35'22.2" S, 137º26'30.1"
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.1, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
