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

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL TO 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("x is 41")
		fallthrough
	case 42:
		fmt.Println("x is 42")
		fallthrough
	case 43:
		fmt.Println("x is 43")
		fallthrough
	default:
		fmt.Println("this is the default case for x")
	}

}