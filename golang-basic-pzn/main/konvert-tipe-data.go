package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 100000
		nilai64       = int64(nilai32)
		nilai8        = int8(nilai32) // batesnya 127
	)

	fmt.Println(nilai32)
	fmt.Println(nilai8) // -96 balik ke titik kerendah (minimum)
	fmt.Println(nilai64)
	
	// konversi byte ke string
	var(
		name = "Adi Nugraha Putra"
		e = name[0] // -> dia akan menjadi byte dulu
		eString = string(e)
	)
	
	fmt.Println(name) // Adi Nugraha Putra
	fmt.Println(eString) // A
}
