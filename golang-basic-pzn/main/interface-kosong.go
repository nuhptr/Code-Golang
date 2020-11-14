package main

import "fmt"

// apapun bisa digunakan
func ups(i int) interface{}  {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data = ups(20)
	fmt.Println(data) // ups
}
