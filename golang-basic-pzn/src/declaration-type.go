package main

import "fmt"

func main() {
	// alias untuk menggantikan sesuatu
	type NoKtp string
	type Married bool

	var(
		noKtpEko NoKtp = "123456789"
		marriedStatus Married =  true
	)

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}