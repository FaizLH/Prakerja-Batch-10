// menentukan nilai terkecil dan terbesar dari nilai yang dimasukan
package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, v := range numbers {
		if *v < min {
			min = *v
		} else if *v > max {
			max = *v
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Print("masukan nilai a1 : ")
	fmt.Scanln(&a1)
	fmt.Print("masukan nilai a2 : ")
	fmt.Scanln(&a2)
	fmt.Print("masukan nilai a3 : ")
	fmt.Scanln(&a3)
	fmt.Print("masukan nilai a4 : ")
	fmt.Scanln(&a4)
	fmt.Print("masukan nilai a5 : ")
	fmt.Scanln(&a5)
	fmt.Print("masukan nilai a6 : ")
	fmt.Scanln(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("nilai min :", min)
	fmt.Println("nilai max :", max)
}
