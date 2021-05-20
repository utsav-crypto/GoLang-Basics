package main

import "fmt"

func main() {
	x := 4
	switch x {
	case 1, 4:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")

	default:
		fmt.Println("Not a case")
	}
}
