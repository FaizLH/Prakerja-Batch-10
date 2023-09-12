// menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill)
package main

import "fmt"

type Car struct {
	tipe   string
	fuelIn float64
}

func (car Car) perkiraanJarak() float64 {
	konsumsi := 1.5
	return car.fuelIn / konsumsi
}

func main() {
	var car Car
	car.perkiraanJarak()

	fmt.Print("masukan tipe mobil : ")
	fmt.Scanln(&car.tipe)
	fmt.Print("masukan angka bahan bakar : ")
	fmt.Scanln(&car.fuelIn)
	fmt.Println("mobil", car.tipe, "dengan bahan bakar", car.fuelIn, "liter diperkirakan dapat menempuh jarak", car.perkiraanJarak(), "mill.")
}
