package main

import "fmt"

func main() {
	// if else
	var age int
	age = 10
	if age > 10  {
		fmt.Println("Kamu Boleh bermain game")
	} else {
		// execute
		fmt.Println("Tidak Boleh bermain Game")
	}

	// elseif
	score := 80
	var grade string

	if score <= 50  {
		grade = "E"
	}else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		// Execute
		grade = "Null"
	}

	fmt.Println(grade)
}

