package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Adi"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
