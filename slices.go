package main

import "fmt"

func main() {
	fmt.Println("........................................")
	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println(cap(arr[1:len(arr)]))
	// ..................................................
	a := make([]int, 5)
	fmt.Println(a)
}
