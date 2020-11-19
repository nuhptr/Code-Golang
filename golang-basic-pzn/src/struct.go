package main

import "fmt"

func main() {
	var eko Customer
	eko.Name = "Adi Nugraha"
	eko.Address = "Gadingrejo"
	eko.Age = 20

	fmt.Println(eko)

	// struct literal
	joko := Customer{
		Name:    "Joko",
		Age:     30,
		Address: "Indonesia",
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia",
		30,
	}
	fmt.Println(budi)

	//	{Adi Nugraha Gadingrejo 20}
	//	{Joko Indonesia 30}
	//	{Budi Indonesia 30}
}

// Seperti membuat tipe data baru
type Customer struct {
	Name, Address string
	Age           int
}
