// ROLLAR COASTER RIDE

package main

import "fmt"

func main() {
	age := 17
	if age >= 18 {
		fmt.Printf("You can Ride alone.")
	} else if age >= 14 {
		fmt.Printf("You can ride with your father.\n Just wait for %d more years to rSide alone", 18-age)
	} else {
		fmt.Printf("You can not ride.")
	}
}
