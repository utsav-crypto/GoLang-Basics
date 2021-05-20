package main

import "fmt"

func main() {
	crkl := Circle{5.46, &Point{4, 5}}
	fmt.Println(crkl)
	fmt.Println(*crkl.center)
}

type Point struct {
	x int
	y int
}
type Circle struct {
	radius float32
	center *Point
}
