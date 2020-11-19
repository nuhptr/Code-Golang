package main

import "fmt"

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}

// jika type of return same you can declare in last
// if different type it firstName int
func getCompleteName() (
	firstName, middleName, lastName string) {
	firstName = "Adi"
	middleName = "Nugraha"
	lastName = "Putra"

	return firstName, middleName, lastName
}
