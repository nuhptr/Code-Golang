package main

import "fmt"

func main() {
	var (
		names [3]string
	)

	// cara memasukan datanya
	names[0] = "Adi"
	names[1] = "Nugraha"
	names[2] = "Putra"

	// cara menyetaknya
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array secara langsung
	values := [3]int{
		90, 80, 80,
	}
	fmt.Println(values) // [90 80 80]
	fmt.Println(len(values)) // 3
}
