package main

import "fmt"

func main() {
	//number := 5
	//number2 := &number
	//
	//fmt.Println(number)
	//fmt.Println(number2)
	//fmt.Println(*number2)
	//
	//*number2 = 10
	//fmt.Println(*number2)
	//fmt.Println(number)

	//var numberA int = 5
	//var numberB *int = &numberA
	//
	//fmt.Println(numberA)
	//fmt.Println(numberB)
	//fmt.Println(*numberB)
	//
	//numberA = 20
	//fmt.Println(numberA)
	//fmt.Println(numberB)
	//fmt.Println(*numberB)

	 number := 5
	 fmt.Println("Alamat memory: ", &number)
	 fmt.Println("Nilai awal: ", number)

	 change(&number, 100)

	 fmt.Println("Nilai Akhir: ", number)
	 fmt.Println("alamat memori: ", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Almat memori: ", old)
	fmt.Println("Didalam fucntion: ", *old)
}
