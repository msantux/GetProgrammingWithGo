package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("toot ", int(l))
}

type starship struct {
	laser
}

type rover string

func (r rover) talk() string {
	return string(r)
}

func main() {
	var t talker = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))

	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	r := rover("whir whir")
	shout(r)
}
