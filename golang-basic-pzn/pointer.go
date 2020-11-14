package main

import "fmt"

func main() {
	// pass by value
	address1 := Address{
		"Pringsewu", "Lampung",
		"Indonesia",
	}
	address2 := address1

	address2.City = "Papua"
	fmt.Println(address1)
	fmt.Println(address2)
	//  you copy the value not the reference
	// 	{Pringsewu Lampung Indonesia}
	//	{Papua Lampung Indonesia}

	// pass by reference
	family1 := Family{"Miky", "Nela",
		"Syifa"}
	family2 := &family1
	/*	{Miky Nela Syifa}
		&{Miky Nela Syifa} */

	//family2 = &Family{"Jugger", "Maghdalena",
	//	"Rieka"} // new data
	//	{Miky Nela Syifa}
	//	&{Jugger Maghdalena Rieka}

	*family2 = Family{"Jugger", "Maghdalena",
		"Rieka"} // primary data is change
	fmt.Println(family1)
	fmt.Println(family2)

}

type Address struct {
	City, Province, Country string
}

type Family struct {
	father, mother, son string
}
