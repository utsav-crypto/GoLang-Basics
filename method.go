package main

import "fmt"

type Student struct {
	name  string
	grade []int
	age   int
}

func main() {
	s1 := Student{"Utsav", []int{87, 98, 85, 99, 89}, 22}
	average1 := s1.getAvaerage()
	fmt.Println(average1)
}

func (s Student) getAvaerage() float32 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float32(sum) / float32(len(s.grade))
}
