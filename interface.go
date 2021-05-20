package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

func main() {
	c1 := Circle{4.5}
	r1 := Rectangle{3, 4}
	// fmt.Println(c1, r1)
	// fmt.Println(c1.area())
	// fmt.Println(r1.area())
	shapes := []shape{c1, r1}
	// This is one method of using interface
	/*for _, shape := range shapes {
		fmt.Println(shape.area())
	}*/
	//This is another
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}

type Circle struct {
	radius float32
}
type Rectangle struct {
	x int
	y int
}

func (c Circle) area() float32 {
	area := float32(math.Pi * c.radius * c.radius)
	return float32(area)
}
func (r Rectangle) area() float32 {
	area := float32(r.x * r.y)
	return area
}
func getArea(s shape) float32 {
	return s.area()
}
