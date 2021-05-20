package main

import "fmt"

func main() {
	arr := []int{34, 12, 56, 34, 46}
	for i, ele := range arr {
		for j := i + 1; j < len(arr); j++ {
			ele2 := arr[j]
			if ele == ele2 {
				fmt.Println(ele)
			}
		}
	}

}
