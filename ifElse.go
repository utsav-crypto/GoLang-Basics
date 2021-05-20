/*ROLLAR COASTER RIDE*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Write your age for verified ride...\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// age := 17
	if age >= 18 {
		fmt.Printf("You can Ride alone.")
	} else if age >= 14 {
		fmt.Printf("You can ride with your father.\nJust wait for %d more years to ride alone", 18-age)
	} else {
		fmt.Printf("You can not ride.")
	}
}
