package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	brandbury := location{lat: -4.5895, long: 137.4417}
	curiosity := brandbury

	curiosity.long += 0.0106

	fmt.Printf("%+v %+v\n", brandbury, curiosity)
}
