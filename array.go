package main

import "fmt"

func main() {
	arr := [3]int{3, 4, 5}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
	// Multidimensional array
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)
}
