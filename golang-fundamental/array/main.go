package main

import "fmt"

func main() {
	// tidak di isi langsung
	/*var language [5]string
	language[0] = "Go"
	language[1] = "Ruby"
	language[2] = "Javascript"
	language[3] = "C#"
	language[4] = "Python"*/

	// di isi langsung
	// ... -> otomatis menghitung array
	language := [...]string{
		"Ruby",
		"Python",
		"Javascript",
		"Go",
		"C++",
	}
	fmt.Println(language)
	var length = len(language)
	fmt.Println(length)

	for index, lang := range language {
		fmt.Println("Index: ", index, "language: ", lang)
	}
}
