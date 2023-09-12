// menghitung string berapa banyak yang sama didalam sebuah slice
package main

import "fmt"

func Mapping(data []string) map[string]int {
	result := map[string]int{}
	for _, v := range data {
		result[v]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	// map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))
	// map[]
}
