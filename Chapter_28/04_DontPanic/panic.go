package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e, "There is no panic!")
		}
	}()

	panic("I forgot my towel.")

}
