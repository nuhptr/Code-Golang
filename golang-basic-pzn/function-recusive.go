package main

import "fmt"

func main() {
	// fmt.Println(factorialLoop(10)) // 3628800
	fmt.Println(factorialRecusive(5)) // 120
}

func factorialRecusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecusive(value-1)
	}
}
