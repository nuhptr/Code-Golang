package main

import "fmt"

func main() {
	// cara ignore return value
	firstName, lastName, _ := getFullName()
	fmt.Print(firstName, lastName)
}

func getFullName() (string, string, string)  {
	return "Adi ", "Nugraha ", "Makan"
}
