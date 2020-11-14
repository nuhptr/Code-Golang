package main

import "fmt"

func main() {
	// cara pertama
	var name string
	var name2 = "Adi Nugraha Putra"
	var age = 30

	// cara kedua (multiple variable)
	var (
		firstName = "Adi"
		lastName = "Nugraha Putra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	fmt.Println(age) // 30

	name = "Eko Kurniwan"
	fmt.Println(name)

	fmt.Println(name2) // Adi Nugraha Putra

	name = "Nugraha Purta"
	fmt.Println(name)

	// bisa menggunakan := pengganti var
	day := "Night"
	fmt.Println(day) // night
}
