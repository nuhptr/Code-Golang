package main

import "fmt"

func main() {
	people := People{"", "Male", 21}
	// tunjukan func ChangeAddress to struct People
	ChangeAdressToIndonesia(&people)

	fmt.Println(people)
}

type People struct {
	Name, Type string
	Age int
}

func ChangeAdressToIndonesia(people *People)  {
	people.Name = "Adi Nugraha Putra"
}