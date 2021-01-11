package main

import "fmt"

func main() {
	nilaiPersegi := Persegi{Sisi: 10}
	fmt.Println(resultLuas(nilaiPersegi))

	nilaiPersegiPanjang := PersegiPanjang{
		Panjang: 10,
		Lebar:   20,
	}
	fmt.Println(resultLuas(nilaiPersegiPanjang))
}

func resultLuas(nilai Luas) int {
	return nilai.hitungLuas()
}
func (persegiPanjang PersegiPanjang) hitungLuas() int {
	return persegiPanjang.Lebar * persegiPanjang.Panjang
}

func (persegi Persegi) hitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type Luas interface {
	hitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

