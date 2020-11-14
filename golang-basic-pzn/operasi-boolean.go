package main

import "fmt"

func main() {
	var(
		nilaiAkhir = 90
		absensi = 80
		lulusNilaiAkhir = nilaiAkhir > 80
		lulusAbsensi = absensi > 80
		lulus = lulusAbsensi && lulusNilaiAkhir
	)

	fmt.Println(lulusNilaiAkhir) // true
	fmt.Println(absensi) // false
	fmt.Println(lulus) // false

	// cara lain
	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}
