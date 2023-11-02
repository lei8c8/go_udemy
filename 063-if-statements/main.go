package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 45
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("Equal to the meaning of life")
	} else {
		fmt.Println("Geater than the meaning of life")
	}

}