package main

import "fmt"

func main() {
	name := "Adi"

	if name == "Adi" {
		fmt.Println("Hello Adi")
	} else if name == "Budi" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi boleh kenalan?")
	}
	// Hello Adi

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
	// nama sudah benar
}
