// menentukan bilangan prima
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	// Menggunakan akar kuadrat sebagai batas untuk iterasi
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int // Ganti dengan bilangan yang Anda inginkan
	fmt.Print("masukan angka : ")
	fmt.Scanln(&n)

	if isPrime(n) {
		fmt.Println(n, " adalah bilangan prima")
	} else {
		fmt.Println(n, " bukan bilangan prima")
	}
}
