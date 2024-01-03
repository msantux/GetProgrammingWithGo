package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1fº, %.1fº)", l.name, l.lat, l.long)
}

type gps struct {
	world       world
	current     location
	destination location
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1fkm to %v", g.distance(), g.destination.description())
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{name: "Bradbury Landing", lat: -4.5895, long: 137.4417}
	elysium := location{name: "Elysium Planitia", lat: 4.5, long: 135.9}

	gps := gps{
		world:       mars,
		current:     bradbury,
		destination: elysium,
	}

	curiosity := rover{gps: gps}

	fmt.Println("Current location:", curiosity.current.description())
	fmt.Println("Destination:", curiosity.destination.description())
	fmt.Println(curiosity.message())
}
