package main

import "fmt"

func main() {

	/* while loop*/

	//x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	/* for loop*/

	for i := 0; i <= 500; i++ {
		if i != 0 && i%3 == 0 && i%7 == 0 && i%9 == 0 {
			fmt.Println(i)
			continue
		}
	}
}
