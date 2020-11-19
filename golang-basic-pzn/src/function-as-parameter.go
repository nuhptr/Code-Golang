package main

import "fmt"

func main() {
	// Hello Adi
	sayHelloWithFilter("Adi", spamFilter)

	filter := spamFilter
	// Hello ...
	sayHelloWithFilter("Anjing", filter)
}

// type as alias
type Filter func(string) string

func sayHelloWithFilter(name string,
	filter Filter) {

	fmt.Println("Hello " + filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
