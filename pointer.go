package main

import "fmt"

func main() {
	a := 10
	var b *int = &a
	*b = 30
	fmt.Println(a)
	//----------------------------------------------
	tochange := "Hello!"
	fmt.Println(tochange)
	changeValue(&tochange)
	fmt.Println(tochange)
}
func changeValue(str *string) {
	*str = "Voila Changed!"
}
