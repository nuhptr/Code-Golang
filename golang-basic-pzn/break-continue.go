package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // diperulangan yang ke 5 berhenti
		}
		fmt.Println("Perulanngan ke", i)
	}

	fmt.Println("==========")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Peruulangan ke", i)
		// bilangan genap ke skip
	}
}
