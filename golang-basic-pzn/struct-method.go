package main

import "fmt"

func main() {
	rully := CustomerActive{Name: "Rully"}
	rully.sayHello()
	// Hello, My Name is Rully
}

func (customerActive CustomerActive) sayHello() {
	fmt.Println("Hello, My Name is", customerActive.Name)
}

type CustomerActive struct {
	Name, Address string
	age           int
}
