// menentukan bilangan kelipatan 7
package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan angka kelipatan 7 : ")
	fmt.Scanln(&n)

	if n%7 == 0 {
		fmt.Println(n, "adalah bilangan kelipatan 7")
	} else {
		fmt.Println(n, "bukan bilangan kelipatan 7")
	}
}
