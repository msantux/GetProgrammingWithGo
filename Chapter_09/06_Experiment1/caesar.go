package main

import "fmt"

func main() {
	quote := "L fdph, L vdz, L frqtxhuhg"

	for _, c := range quote {
		if (c > 'c' && c <= 'z') || (c > 'C' && c <= 'Z') {
			c -= 3
		} else if (c >= 'a' && c <= 'c') || (c >= 'A' && c <= 'C') {
			c += 23
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
