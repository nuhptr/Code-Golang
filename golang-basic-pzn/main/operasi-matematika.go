package main

import "fmt"

func main() {
	var(
		result = 10 + 10
		a = 10
		b = 10
		c = a + b
		d = a * b
	)
	fmt.Println(result)

	fmt.Println(c) // 20
	fmt.Println(d) // 100

	// augmented assignment
	var i = 10
	i += 10

	fmt.Println(i) // 20

	i++
	fmt.Println(i) // 21
}