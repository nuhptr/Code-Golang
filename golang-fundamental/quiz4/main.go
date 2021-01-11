package main

import (
	"errors"
	"fmt"
)

func main() {
	// passed
	/* scores := []int {1,2,3,4,5}
	total := sum(scores)
	println(total) */

	// passed
	result, err := calculate(10,2, "*")
	if err != nil {
		fmt.Println("Terjadi Error")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func sum(score []int) int {
	var total int
	for _,value := range score {
		total = total + value
	}
	return total
}

func calculate(numberOne , numberTwo int, operator string) (int, error) {
	var result int
	var errorResult error

	switch operator {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		result = numberOne / numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}

	return result, errorResult
}