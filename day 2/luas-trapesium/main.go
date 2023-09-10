package main

import "fmt"

func main() {
	var a, b, t int
	fmt.Print("masukan nilai sisi sejajar yang lebih pendek : ")
	fmt.Scanln(&a)
	fmt.Print("masukan nilai sisi sejajar yang lebih panjang : ")
	fmt.Scanln(&b)
	fmt.Print("masukan nilai tinggi : ")
	fmt.Scanln(&t)

	L := (a + b) * t / 2

	fmt.Println(L)
}
