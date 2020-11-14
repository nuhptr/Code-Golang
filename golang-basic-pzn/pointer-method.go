package main

import "fmt"

func main() {
	eko := Man{"eko"}
	eko.Married()
	fmt.Println(eko)
}

type Man struct {
	Name string
}

func (man *Man) Married()  {
	man.Name = "Mr. " + man.Name
}


