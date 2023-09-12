// menggabungkan 2 array tapi jangan sampai terdapat nama yang sama
package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// yoyr code here
	var result []string
	var mapResult = map[string]bool{}

	arrayA = append(arrayA, arrayB...)
	isFound := false

	for _, v := range arrayA {
		_, isFound = mapResult[v]
		if isFound == false {
			result = append(result, v)
		}
		mapResult[v] = true
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "Law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

}
