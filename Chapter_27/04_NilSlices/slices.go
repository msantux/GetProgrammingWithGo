package main

import "fmt"

func main() {
	var soup []string

	fmt.Println(soup == nil)

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))

	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)

	soup = mirepoix(nil)
	fmt.Println(soup)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}
