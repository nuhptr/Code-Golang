package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := Pembagian(100,0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi dengan 0")
	} else {
		return nilai / pembagi, nil
	}
}
