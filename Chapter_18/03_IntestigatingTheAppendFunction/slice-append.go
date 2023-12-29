package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}
	dump("dwarfs", dwarfs)

	dwarfs2 := append(dwarfs, "Orcus")
	dump("dwarfs2", dwarfs2)

	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)
}
