package main

import "fmt"

func main() {
	printMyResult("Makan")
	fmt.Println(printMyResult("Minum "))
	fmt.Println("================")

	fmt.Println(add(10,2))

	luas, keliling := calculate(10,2)
	fmt.Println(luas) // 20
	fmt.Println(keliling) // 24
}

func printMyResult(word string) string{
	fmt.Println("Saya sedang belajar", word)
	newSentence := word + "sambil berdoa"
	return newSentence
}

func add(numberOne int, numberTwo int) int {
	total := numberOne + numberTwo
	return total
}

// function Multiple Return
func calculate(panjang int, lebar int) (int, int)  {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}