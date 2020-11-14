package main

import "fmt"

func main() {
	var result = random()
	// assertion manual
	//resultString := result.(string)
	//fmt.Println(resultString)

	// dengan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return 100
}
