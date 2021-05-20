package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  1,
		"pear":   2,
		"orange": 3,
	}
	fmt.Println(mp)
	// second method to create map using make
	mp2 := make(map[string]int)
	mp2 = map[string]int{
		"utsav":    1,
		"hemant":   2,
		"somendra": 3,
	}
	fmt.Println(mp2)
	mp2["somendra"] = 45
	delete(mp2, "somendra")
	fmt.Println(mp2)

	// checking of key

	val, ok := mp["apple"]
	fmt.Println(val, ok)
}
