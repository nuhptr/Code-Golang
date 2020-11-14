package main

import "fmt"

func main() {
	result := getHello("Adi Nugraha Putra")
	result2 := getHello("")
	fmt.Println(result)
	fmt.Println(result2) // hello bro
}

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}
