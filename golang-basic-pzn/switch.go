package main

import "fmt"

func main() {
	name := "Adi"

	switch name {
	case "Adi":
		fmt.Println("Hello Adi")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}
	// Hello Adi

	name2 := "nugraha"
	// switch short condition
	switch length := len(name2); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	} // nama terlalu panjang
}
