package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Makan": 10,
		"Minum": 11,
	}

	for key, value := range myMap {
		fmt.Println("Key:", key, "value:", value)
	}

	fmt.Println("==========")

	delete(myMap, "Makan")

	for key, value := range myMap {
		fmt.Println("Key:", key, "value:", value)

	}

	_, isAvailable := myMap["python" ]
	fmt.Println(isAvailable) //false
}