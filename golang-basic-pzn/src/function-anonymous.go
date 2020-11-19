package main

import "fmt"

func main() {
	// insert to variable
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registersUser("anjing", blacklist)

	// insert in other function
	registersUser("adi", func(name string) bool {
		return name == "anjing"
	})
}

type Blacklist func(string) bool

func registersUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
