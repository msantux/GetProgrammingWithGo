package main

import (
	"fmt"

	"santux.com/GetProgrammingWithGo/GetProgrammingWithGo/GeeksForGeeks/02_IdentifiersGoLanguage/file2"
)

var ExportedVariable = "Hello, World!"

func main() {
	fmt.Println(ExportedVariable)

	fmt.Println(file2.AnotherExportedVariable)
}
