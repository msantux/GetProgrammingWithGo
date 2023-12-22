package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		fmt.Print("Offset ", offset, ": ")
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())

	offset = 10
	fmt.Println(sensor())

	sensor = calibrate(fakeSensor, offset)
	fmt.Println("============Fake Sensor============")
	for i := 0; i < 10; i++ {
		fmt.Println(sensor())
		time.Sleep(time.Second)
	}
}
