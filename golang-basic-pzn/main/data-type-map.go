package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Adi",
		"address": "Gadingrejo",
	}

	/* menambahkan
	person["hobby"] = "makan"
	*/

	/*
	delete(person, "hobby")
	*/

	fmt.Println(person) // map[address:Gadingrejo name:Adi]
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
