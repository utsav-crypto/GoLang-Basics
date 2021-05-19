package main

import "fmt"

func main() {
	var num1 = float32(45)
	var num2 = 23.67
	result := num1 + float32(num2)
	fmt.Println(".........................................................")
	fmt.Printf("Answer: %f", result)
	a := 5
	b := 6
	res := a < b
	fmt.Println("\n.........................................................")
	fmt.Printf("Answer would be %t", res)
	out := a > b || a < 6
	fmt.Println("\n.........................................................")
	fmt.Printf("This is new answer which is also %t", out)
	val := (true || false) && true
	fmt.Println("\n.........................................................")
	fmt.Printf("%t", val)

}
