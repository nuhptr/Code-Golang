package main

import "fmt"

func main() {
	student := []map[string] string {
		{"name": "Agung", "score" : "A"},
		{"name" : "Adi", "score" : "B"},
	}

	for _, students := range student{
		fmt.Println(students["name"], "-", students["score"])
	}
}
