package main

import "fmt"

func main() {
	// for i := 0; i<= 100; i++ {
	//	fmt.Println("saya sedang belajar go : ", i)
	// }

	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index: ", index, "letter: ", string(letter))
	}
}
