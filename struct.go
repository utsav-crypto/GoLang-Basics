package main

import "fmt"

func main() {
	u1 := Utsav{1, 2}
	u1.x = 34
	fmt.Println(u1.x)
	u2 := Utsav{y: 4}
	fmt.Println(u2)
	u3 := &Utsav{y: 400}
	changeX(u3)
	fmt.Println(*u3)
}

type Utsav struct {
	x int
	y int
}

func changeX(pt *Utsav) {
	pt.x = 300
}
