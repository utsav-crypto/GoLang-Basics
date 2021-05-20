package main

import "fmt"

func test2(myfunc func(int) int) {
	fmt.Println(myfunc(7))
}

func main() {
	x := func(x int) {
		fmt.Println("Hello!", x)
	}
	x(5)

	// passing inline arguments
	y := func(y int) int {
		return y * -1
	}(8)
	fmt.Println(y)
	//passing function as arguments
	test := func(x int) int {
		return x * -1
	}
	test2(test)
}
