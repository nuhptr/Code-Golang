package main

import "fmt"

func main() {
	runApplication(1)
	runApp(true) // panic: error woi!
}

func logging() {
	fmt.Println("Seleasai memanggil function")
}

func runApplication(value int) {
	defer logging() // membuat sebuah function jalan
	// diakhir walaupun error
	fmt.Println("run application")
	result := 10 / value
	// panic: runtime error: integer divide by zero
	fmt.Println("result", result)
}

func endApp() {
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("error woi!")
	}
}
