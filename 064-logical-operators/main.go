package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

}