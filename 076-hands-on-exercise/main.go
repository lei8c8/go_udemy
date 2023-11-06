package main

import (
	"fmt"
	"math/rand"
)

var g int

func init() {
	fmt.Println("This is where the init function in effective go")
	g = 100
	fmt.Println(g)
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\t", x)
	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 100 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}
}
