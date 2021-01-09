package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, Dunia Golang")
	// Halo, Dunia Golang

	// sentence := TestAja()
	// fmt.Println(sentence)
	// Halo nama saya Adi dan umur saya 20

	result := calculation.Add(20, 20)
	fmt.Println(result) // 40

	result2 := calculation.Add(12, 20)
	fmt.Println(result2) // 32
}
