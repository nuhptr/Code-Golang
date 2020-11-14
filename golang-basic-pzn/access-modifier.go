package main

import "fmt"

func main() {
	// tidak bisa diakses dari luar
	name := "ADI"

	// bisa diakses dari luar package
	var Name2 string
	fmt.Println(name)
	fmt.Println(Name2)
}

// bisa diakses dari luar
func Run()  {

}

// tidak bisa diakses dari luar
func handsome()  {

}


