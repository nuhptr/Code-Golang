package main

import "fmt"

func main() {
	var eko Person
	eko.Name = "Eko"
	sayHelloPerson(eko)

	animal := Animal{Name: "Kucing"}
	sayHelloPerson(animal)
	// 	Hello Eko
	//	Hello Kucing
}

type HasName interface {
	GetName() string
}

func sayHelloPerson(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// struct pertama
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// struct kedua
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
