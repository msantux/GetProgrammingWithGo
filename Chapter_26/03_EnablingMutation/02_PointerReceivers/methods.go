package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan)

	const layout = "Mon, Jan 2, 2006"

	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))
}
