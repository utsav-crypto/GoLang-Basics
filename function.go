package main

import "fmt"

func add(x, y int) (int, int) {
	defer fmt.Println("first defer")
	return x + y, x - y
}

// lebelling the return type
func add1(x, y int) (z1, z2 int) {
	defer fmt.Println("second defer")
	z1 = x + y
	z2 = y - x
	return
}
func main() {
	defer fmt.Println("Coding is really fun.")
	ans1, ans2 := add(4, 7)
	answer1, answer2 := add1(5, 8)
	fmt.Println(ans1, ans2)
	fmt.Println(answer1, answer2)
}
