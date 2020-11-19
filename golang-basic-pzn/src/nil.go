package main

import "fmt"

func main() {
	person := newMap("Adi")
	fmt.Println(person) // map[name:Adi]

	data := newMap("") // data kosong
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}
}
