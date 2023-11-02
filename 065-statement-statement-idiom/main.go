package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EUQAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
	
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EUQAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}