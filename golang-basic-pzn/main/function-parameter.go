package main

import "fmt"

func main() {
	sayHelloTo("Adi", "Nugraha")
}

func sayHelloTo(firstName string, lastName string)  {
	fmt.Println("Hello", firstName, lastName)
}