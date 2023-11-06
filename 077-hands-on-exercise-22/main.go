package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x is %v and y is %v\n", x, y)
	
	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is from 4 to 6 inclusive of both numbers")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the cases")
	}
}