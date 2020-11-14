package main

import "fmt"

func main() {
	names := [...]string{ // array
		"Adi", "Nugraha", "Putra",
		"mau", "Makan", "Sekarang",
	}

	slice := names[4:6] // ngambil data ke 4 - sebelum 6
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(cap(slice)) // 2

	// cara paling aman
	newSlice := make([]string, 2, 5) // slice
	newSlice[0] = "Adi"
	newSlice[1] = "Nugraha"

	fmt.Println(newSlice)
}
