package main

import (
	"fmt"
	"quiz/multiply"
)

func main() {
	multiplyResult := multiply.Multiply(100,100)
	fmt.Println(multiplyResult) // 10000 -> done
}
