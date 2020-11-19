package main

import "fmt"

func main() {
	counter := 0

	// bisa akses data diatasnya
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	// gabisa akses data di dalam anonymous func
	increment()
	increment()
	fmt.Println(counter)
}
