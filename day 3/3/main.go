// membuat array yang berisi angka yang hanya muncul sekali pada string
package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	// your code here
	result := []int{}
	mapResult := map[string]int{}

	for _, v := range angka {
		mapResult[string(v)]++
	}

	var n int
	for k, v := range mapResult {
		if v == 1 {
			n, _ = strconv.Atoi(k)
			result = append(result, n)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
