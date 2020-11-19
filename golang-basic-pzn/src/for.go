package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	// melakukan perulangan sampai 10

	// for dengan statment
	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan ke ", counter1)
	}

	// for range
	slice := []string{"Adi", "Nugraha", "Putra"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// Adi Nugraha Putra

	names := map[string]string{
		"aku":  "siapa",
		"kamu": "siapanya dia?",
	}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	// index aku = siapa
	// index kamu = siapanya dia?
}
