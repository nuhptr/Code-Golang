package main

import "fmt"

func main() {
	total := sumAll(10,20,20,30,30,40,40)
	fmt.Println(total) // 190

	slice := []int {10,10,10,10}
	total = sumAll(slice...) // slice -> variadic
	fmt.Println(total) // 40
}

/*create variadic argument can only one in function
  parameter and cannot put in first, middle, only last */

func sumAll(numbers ...int) int {
	total := 0
	// _ -> boleh dikosongkan
	for _, number := range numbers {
		total += number
	}

	return total
}