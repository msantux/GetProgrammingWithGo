package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	plannetsMarkII := planets

	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(plannetsMarkII)
}
