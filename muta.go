package main

import "fmt"

func main() {
	// mutable->slice,map,channel,array
	// #slice
	var slice []int = []int{2, 3, 4, 5}
	slice2 := slice
	slice2[2] = 35
	fmt.Println(slice2)
	// #array
	array := [5]int{45, 67, 34, 29, 83}
	array[3] = 47
	fmt.Println(array)
	// #maps
	mp := map[string]int{
		"a": 1,
		"b": 2,
	}
	mp["b"] = 14
	fmt.Println(mp)
}
